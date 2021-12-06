package comutil

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"log"
)

//采样所有追踪（不能再online环境使用）
const JaegerSamplerParam = 1
const JaegerReportingHost = "192.168.59.131:6831"

//将GlobalTracerHandler作为全局变量使用，这样保证代码中使用同一个tracer
var GlobalTracerHandler *TraceHandler

//初始化服务
type TraceHandler struct {
	Tracer opentracing.Tracer
	Closer io.Closer
}

func init() {
	GlobalTracerHandler = InitTracer()
}

//jaeger 初始化
func InitTracer() *TraceHandler {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: JaegerSamplerParam,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: JaegerReportingHost,
		},
	}
	//设置服务名称
	cfg.ServiceName = "jaeger_test"
	//创建tracer
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		//失败错误
		log.Fatal(err)
	}
	return &TraceHandler{
		Tracer: tracer,
		Closer: closer,
	}
}
