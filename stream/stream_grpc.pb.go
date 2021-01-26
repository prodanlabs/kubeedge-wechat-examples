// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stream

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StreamClientClient is the client API for StreamClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamClientClient interface {
	// 简单模式 rpc ，客户端通过stub发起请求,等待服务端返回结果;
	SimpleMode(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// 客户端流式 rpc ，服务端等到客户端所有请求发送完毕后，向客户端发送一次消息;
	Upload(ctx context.Context, opts ...grpc.CallOption) (StreamClient_UploadClient, error)
}

type streamClientClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamClientClient(cc grpc.ClientConnInterface) StreamClientClient {
	return &streamClientClient{cc}
}

func (c *streamClientClient) SimpleMode(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/stream.StreamClient/SimpleMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamClientClient) Upload(ctx context.Context, opts ...grpc.CallOption) (StreamClient_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamClient_ServiceDesc.Streams[0], "/stream.StreamClient/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamClientUploadClient{stream}
	return x, nil
}

type StreamClient_UploadClient interface {
	Send(*StreamRequest) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type streamClientUploadClient struct {
	grpc.ClientStream
}

func (x *streamClientUploadClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamClientUploadClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamClientServer is the server API for StreamClient service.
// All implementations must embed UnimplementedStreamClientServer
// for forward compatibility
type StreamClientServer interface {
	// 简单模式 rpc ，客户端通过stub发起请求,等待服务端返回结果;
	SimpleMode(context.Context, *Request) (*Response, error)
	// 客户端流式 rpc ，服务端等到客户端所有请求发送完毕后，向客户端发送一次消息;
	Upload(StreamClient_UploadServer) error
	mustEmbedUnimplementedStreamClientServer()
}

// UnimplementedStreamClientServer must be embedded to have forward compatible implementations.
type UnimplementedStreamClientServer struct {
}

func (UnimplementedStreamClientServer) SimpleMode(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SimpleMode not implemented")
}
func (UnimplementedStreamClientServer) Upload(StreamClient_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedStreamClientServer) mustEmbedUnimplementedStreamClientServer() {}

// UnsafeStreamClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamClientServer will
// result in compilation errors.
type UnsafeStreamClientServer interface {
	mustEmbedUnimplementedStreamClientServer()
}

func RegisterStreamClientServer(s grpc.ServiceRegistrar, srv StreamClientServer) {
	s.RegisterService(&StreamClient_ServiceDesc, srv)
}

func _StreamClient_SimpleMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamClientServer).SimpleMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream.StreamClient/SimpleMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamClientServer).SimpleMode(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamClient_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamClientServer).Upload(&streamClientUploadServer{stream})
}

type StreamClient_UploadServer interface {
	SendAndClose(*Response) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamClientUploadServer struct {
	grpc.ServerStream
}

func (x *streamClientUploadServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamClientUploadServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamClient_ServiceDesc is the grpc.ServiceDesc for StreamClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream.StreamClient",
	HandlerType: (*StreamClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SimpleMode",
			Handler:    _StreamClient_SimpleMode_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _StreamClient_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "stream/stream.proto",
}
