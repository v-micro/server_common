package serverorder

import (
	"context"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"sync"
	"time"
)

type Client struct {
	lock     sync.Mutex
	Ctx      context.Context
	Conn     *grpc.ClientConn

	//应有的服务
	Ping PingClient
}

//公共方法
var grpcUrl = "127.0.0.1:10001"
var grpcClient Client

//初始化
func init() {
	grpcClient = Client{
		Ctx: context.Background(),
	}
}

//实例化
func GetClient() *Client {
	var err error

	//判断是否存在
	if grpcClient.Conn != nil && grpcClient.Conn.GetState().String() != "SHUTDOWN" {
		return &grpcClient
	}

	//互斥锁
	grpcClient.lock.Lock()
	defer grpcClient.lock.Unlock()

	//判断是否存在
	if grpcClient.Conn != nil && grpcClient.Conn.GetState().String() != "SHUTDOWN" {
		return &grpcClient
	}

	//设置超时
	grpcClient.Ctx, _ = context.WithTimeout(context.Background(), 5*time.Minute)

	if grpcClient.Conn, err = grpc.Dial(grpcUrl, grpc.WithInsecure()); err != nil {
		glog.Error("grpc连接失败，", err , grpcUrl)
		return nil
	} else {
		//rpc链接服务
		grpcClient.Ping = NewPingClient(grpcClient.Conn)

		return &grpcClient
	}
}
