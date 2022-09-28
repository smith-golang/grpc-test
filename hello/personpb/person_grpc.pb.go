// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: personpb/person.proto

package personpb

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PersonServiceClient is the client API for PersonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonServiceClient interface {
}

type personServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonServiceClient(cc grpc.ClientConnInterface) PersonServiceClient {
	return &personServiceClient{cc}
}

// PersonServiceServer is the server API for PersonService service.
// All implementations must embed UnimplementedPersonServiceServer
// for forward compatibility
type PersonServiceServer interface {
	mustEmbedUnimplementedPersonServiceServer()
}

// UnimplementedPersonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPersonServiceServer struct {
}

func (UnimplementedPersonServiceServer) mustEmbedUnimplementedPersonServiceServer() {}

// UnsafePersonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonServiceServer will
// result in compilation errors.
type UnsafePersonServiceServer interface {
	mustEmbedUnimplementedPersonServiceServer()
}

func RegisterPersonServiceServer(s grpc.ServiceRegistrar, srv PersonServiceServer) {
	s.RegisterService(&PersonService_ServiceDesc, srv)
}

// PersonService_ServiceDesc is the grpc.ServiceDesc for PersonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PersonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "person.PersonService",
	HandlerType: (*PersonServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "personpb/person.proto",
}
