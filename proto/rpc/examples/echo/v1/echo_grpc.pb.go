// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoServiceClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	EchoMultiple(ctx context.Context, in *EchoMultipleRequest, opts ...grpc.CallOption) (EchoService_EchoMultipleClient, error)
	EchoBiDi(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoBiDiClient, error)
}

type echoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoServiceClient(cc grpc.ClientConnInterface) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/proto.rpc.examples.echo.v1.EchoService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoMultiple(ctx context.Context, in *EchoMultipleRequest, opts ...grpc.CallOption) (EchoService_EchoMultipleClient, error) {
	stream, err := c.cc.NewStream(ctx, &EchoService_ServiceDesc.Streams[0], "/proto.rpc.examples.echo.v1.EchoService/EchoMultiple", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceEchoMultipleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EchoService_EchoMultipleClient interface {
	Recv() (*EchoMultipleResponse, error)
	grpc.ClientStream
}

type echoServiceEchoMultipleClient struct {
	grpc.ClientStream
}

func (x *echoServiceEchoMultipleClient) Recv() (*EchoMultipleResponse, error) {
	m := new(EchoMultipleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoServiceClient) EchoBiDi(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoBiDiClient, error) {
	stream, err := c.cc.NewStream(ctx, &EchoService_ServiceDesc.Streams[1], "/proto.rpc.examples.echo.v1.EchoService/EchoBiDi", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceEchoBiDiClient{stream}
	return x, nil
}

type EchoService_EchoBiDiClient interface {
	Send(*EchoBiDiRequest) error
	Recv() (*EchoBiDiResponse, error)
	grpc.ClientStream
}

type echoServiceEchoBiDiClient struct {
	grpc.ClientStream
}

func (x *echoServiceEchoBiDiClient) Send(m *EchoBiDiRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoServiceEchoBiDiClient) Recv() (*EchoBiDiResponse, error) {
	m := new(EchoBiDiResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoServiceServer is the server API for EchoService service.
// All implementations must embed UnimplementedEchoServiceServer
// for forward compatibility
type EchoServiceServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	EchoMultiple(*EchoMultipleRequest, EchoService_EchoMultipleServer) error
	EchoBiDi(EchoService_EchoBiDiServer) error
	mustEmbedUnimplementedEchoServiceServer()
}

// UnimplementedEchoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEchoServiceServer struct {
}

func (UnimplementedEchoServiceServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedEchoServiceServer) EchoMultiple(*EchoMultipleRequest, EchoService_EchoMultipleServer) error {
	return status.Errorf(codes.Unimplemented, "method EchoMultiple not implemented")
}
func (UnimplementedEchoServiceServer) EchoBiDi(EchoService_EchoBiDiServer) error {
	return status.Errorf(codes.Unimplemented, "method EchoBiDi not implemented")
}
func (UnimplementedEchoServiceServer) mustEmbedUnimplementedEchoServiceServer() {}

// UnsafeEchoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EchoServiceServer will
// result in compilation errors.
type UnsafeEchoServiceServer interface {
	mustEmbedUnimplementedEchoServiceServer()
}

func RegisterEchoServiceServer(s grpc.ServiceRegistrar, srv EchoServiceServer) {
	s.RegisterService(&EchoService_ServiceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.rpc.examples.echo.v1.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoMultiple_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EchoMultipleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EchoServiceServer).EchoMultiple(m, &echoServiceEchoMultipleServer{stream})
}

type EchoService_EchoMultipleServer interface {
	Send(*EchoMultipleResponse) error
	grpc.ServerStream
}

type echoServiceEchoMultipleServer struct {
	grpc.ServerStream
}

func (x *echoServiceEchoMultipleServer) Send(m *EchoMultipleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EchoService_EchoBiDi_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServiceServer).EchoBiDi(&echoServiceEchoBiDiServer{stream})
}

type EchoService_EchoBiDiServer interface {
	Send(*EchoBiDiResponse) error
	Recv() (*EchoBiDiRequest, error)
	grpc.ServerStream
}

type echoServiceEchoBiDiServer struct {
	grpc.ServerStream
}

func (x *echoServiceEchoBiDiServer) Send(m *EchoBiDiResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoServiceEchoBiDiServer) Recv() (*EchoBiDiRequest, error) {
	m := new(EchoBiDiRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoService_ServiceDesc is the grpc.ServiceDesc for EchoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EchoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.rpc.examples.echo.v1.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EchoMultiple",
			Handler:       _EchoService_EchoMultiple_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EchoBiDi",
			Handler:       _EchoService_EchoBiDi_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/rpc/examples/echo/v1/echo.proto",
}
