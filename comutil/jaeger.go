package comutil

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

// 初始化jaeger
func InitJaeger(service_name string,host_port string) (tracer opentracing.Tracer, closer io.Closer, err error) {
	// 构造配置信息
	cfg := &config.Configuration{
		// 设置服务名称
		ServiceName: service_name,
		// 设置采样参数
		Sampler: &config.SamplerConfig{
			Type:  "const", // 全采样模式
			Param: 1,       // 开启全采样模式
		},
		Reporter: &config.ReporterConfig{
			LogSpans           : true,
			LocalAgentHostPort : host_port,
		},
	}
	// 生成一条新tracer
	tracer, closer, err = cfg.NewTracer()
	if err == nil {
		// 设置tracer为全局单例对象
		opentracing.SetGlobalTracer(tracer)
	}
	return
}
