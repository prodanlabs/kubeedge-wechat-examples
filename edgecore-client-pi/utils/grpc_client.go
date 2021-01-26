package utils

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/prodanlabs/kubeedge-wechat-examples/stream"
)

// PUTMenu 向服务端上传结果
func PUTMenu() {

	// 连接服务器
	// conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	// Set up a connection to the server.
	cred, err := credentials.NewClientTLSFromFile("server.crt", "xxx")
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}

	conn, err := grpc.Dial("xxx:443", grpc.WithTransportCredentials(cred), grpc.WithBlock())
	if err != nil {
		log.Println("Failed to connect to grpc server", err)
	}
	defer conn.Close()
	// 建立 gRPC 连接
	streamClient := pb.NewStreamClientClient(conn)
	route(streamClient)
	musicList(streamClient)
}

// route 调用服务端 SimpleMode 方法
func route(streamClient pb.StreamClientClient) {
	// 创建发送结构体
	req := pb.Request{
		Data: "grpc",
	}
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := streamClient.SimpleMode(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Printf("status: %d %s\n", res.Code, res.Value)
}

// musicList 调用服务端 Upload 方法
func musicList(streamClient pb.StreamClientClient) {
	// 调用服务端RouteList方法，获流
	stream, err := streamClient.Upload(context.Background())
	if err != nil {
		log.Fatalf("Upload list err: %v", err)
	}

	mp3 := getFile("/home/pi/music/")
	for i := range mp3 {
		// 向流中发送消息
		file := formatFileName(mp3[i])
		err := stream.Send(&pb.StreamRequest{Mun: int32(i + 1), StreamData: file})
		// 发送也要检测EOF，当服务端在消息没接收完前主动调用SendAndClose()关闭stream，此时客户端还执行Send()，则会返回EOF错误，所以这里需要加上io.EOF判断
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
	}
	//关闭流并获取返回的消息
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("RouteList get response err: %v", err)
		return
	}
	log.Printf("status: Server Receiving %s!", res.Value)
}
