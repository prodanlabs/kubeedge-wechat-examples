package handle

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"google.golang.org/grpc"

	pb "github.com/prodanlabs/kubeedge-wechat-examples/stream"
)

var S *StreamServer

func init() {
	S = &StreamServer{}
	S.Ch = make(chan string, 1)
	S.Songs = ""
}

type StreamServer struct {
	pb.UnsafeStreamClientServer
	Ch    chan string
	Songs string
}

// 实现 SimpleMode 方法
func (s *StreamServer) SimpleMode(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	res := pb.Response{
		Code:  200,
		Value: req.Data + " ok!",
	}
	return &res, nil
}

// 实现 Upload 方法
func (s *StreamServer) Upload(scu pb.StreamClient_UploadServer) error {
	for {
		// 从流中获取消息
		res, err := scu.Recv()
		if err == io.EOF {
			// 向通道发送关闭信号
			S.Ch <- "EOF"
			log.Println("Server receiving completed")
			return scu.SendAndClose(&pb.Response{Value: "Success"})
		}
		if err != nil {
			return err
		}
		// *可以直接存储在 StreamServer 结构体中 ，这里是为了测试 chan 通道
		// 获取的消息通过通道传递给 chan
		S.Ch <- strconv.Itoa(int(res.Mun)) + ": " + res.StreamData
	}
}

// RegisterGRPC gRPC server.
func RegisterGRPC(app *iris.Application) {
	// 实例化 GRPC 对象
	grpcServer := grpc.NewServer()
	ctrl := &StreamServer{}
	pb.RegisterStreamClientServer(grpcServer, ctrl)

	// Register MVC application controller for gRPC services.
	serviceName := fmt.Sprintf("%s", pb.File_stream_stream_proto.Services().Get(0).FullName())
	mvc.New(app).Handle(ctrl, mvc.GRPC{
		Server:      grpcServer,  // Required.
		ServiceName: serviceName, // Required. 服务的名称必须是用于构建 gRPC 路由的路径
		Strict:      true,        // 如果Strict选项为true，则此控制器将仅服务于基于gRPC的客户端，在普通HTTP客户机上触发404
	})
}
