// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: bidir_stream/bidir_streampb/bidir_streampb.proto

package bidir_streampb

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

// GreetingServiceClient is the client API for GreetingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetingServiceClient interface {
	// Bi Directional Straming
	GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetingService_GreetEveryoneClient, error)
}

type greetingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetingServiceClient(cc grpc.ClientConnInterface) GreetingServiceClient {
	return &greetingServiceClient{cc}
}

func (c *greetingServiceClient) GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetingService_GreetEveryoneClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetingService_ServiceDesc.Streams[0], "/bidir_stream.GreetingService/GreetEveryone", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetingServiceGreetEveryoneClient{stream}
	return x, nil
}

type GreetingService_GreetEveryoneClient interface {
	Send(*GreetEveryoneRequest) error
	Recv() (*GreetEveryoneResponse, error)
	grpc.ClientStream
}

type greetingServiceGreetEveryoneClient struct {
	grpc.ClientStream
}

func (x *greetingServiceGreetEveryoneClient) Send(m *GreetEveryoneRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetingServiceGreetEveryoneClient) Recv() (*GreetEveryoneResponse, error) {
	m := new(GreetEveryoneResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetingServiceServer is the server API for GreetingService service.
// All implementations must embed UnimplementedGreetingServiceServer
// for forward compatibility
type GreetingServiceServer interface {
	// Bi Directional Straming
	GreetEveryone(GreetingService_GreetEveryoneServer) error
	mustEmbedUnimplementedGreetingServiceServer()
}

// UnimplementedGreetingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetingServiceServer struct {
}

func (UnimplementedGreetingServiceServer) GreetEveryone(GreetingService_GreetEveryoneServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetEveryone not implemented")
}
func (UnimplementedGreetingServiceServer) mustEmbedUnimplementedGreetingServiceServer() {}

// UnsafeGreetingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetingServiceServer will
// result in compilation errors.
type UnsafeGreetingServiceServer interface {
	mustEmbedUnimplementedGreetingServiceServer()
}

func RegisterGreetingServiceServer(s grpc.ServiceRegistrar, srv GreetingServiceServer) {
	s.RegisterService(&GreetingService_ServiceDesc, srv)
}

func _GreetingService_GreetEveryone_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetingServiceServer).GreetEveryone(&greetingServiceGreetEveryoneServer{stream})
}

type GreetingService_GreetEveryoneServer interface {
	Send(*GreetEveryoneResponse) error
	Recv() (*GreetEveryoneRequest, error)
	grpc.ServerStream
}

type greetingServiceGreetEveryoneServer struct {
	grpc.ServerStream
}

func (x *greetingServiceGreetEveryoneServer) Send(m *GreetEveryoneResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetingServiceGreetEveryoneServer) Recv() (*GreetEveryoneRequest, error) {
	m := new(GreetEveryoneRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetingService_ServiceDesc is the grpc.ServiceDesc for GreetingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bidir_stream.GreetingService",
	HandlerType: (*GreetingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetEveryone",
			Handler:       _GreetingService_GreetEveryone_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bidir_stream/bidir_streampb/bidir_streampb.proto",
}