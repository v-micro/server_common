package serverorder

import (
	"context"
	"github.com/golang/glog"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"time"
)

type Client struct {
	Ctx      context.Context
	Conn     *grpc.ClientConn
	//应有的服务
	Ping PingClient
}

//公共方法
var grpcClient Client

//初始化
func init() {
	grpcClient = Client{
		Ctx: context.Background(),
	}
}

//实例化
func GetClient(grpcUrl string,tracer opentracing.Tracer) *Client {
	var err error

	//grpcUrl 判断
	if grpcUrl == "" {
		glog.Error("grpcUrl 连接获取失败")
		return nil
	}

	//设置超时
	grpcClient.Ctx, _ = context.WithTimeout(context.Background(), 5*time.Minute)

	//是否使用链路
	if tracer != nil {
		grpcClient.Conn, err = grpc.Dial(grpcUrl, grpc.WithInsecure(),grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(tracer))));
	}else{
		grpcClient.Conn, err = grpc.Dial(grpcUrl, grpc.WithInsecure());
	}

	//判断是否grpc成功
	if err != nil {
		glog.Error("grpc连接失败，", err , grpcUrl)
		return nil
	}

	//rpc链接服务
	grpcClient.Ping = NewPingClient(grpcClient.Conn)

	return &grpcClient
}
