// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: counter.proto

package counter

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

// IncrementCounterClient is the client API for IncrementCounter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IncrementCounterClient interface {
	Increment(ctx context.Context, in *IncrementBy, opts ...grpc.CallOption) (*Status, error)
}

type incrementCounterClient struct {
	cc grpc.ClientConnInterface
}

func NewIncrementCounterClient(cc grpc.ClientConnInterface) IncrementCounterClient {
	return &incrementCounterClient{cc}
}

func (c *incrementCounterClient) Increment(ctx context.Context, in *IncrementBy, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/counter.incrementCounter/increment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IncrementCounterServer is the server API for IncrementCounter service.
// All implementations must embed UnimplementedIncrementCounterServer
// for forward compatibility
type IncrementCounterServer interface {
	Increment(context.Context, *IncrementBy) (*Status, error)
	mustEmbedUnimplementedIncrementCounterServer()
}

// UnimplementedIncrementCounterServer must be embedded to have forward compatible implementations.
type UnimplementedIncrementCounterServer struct {
}

func (UnimplementedIncrementCounterServer) Increment(context.Context, *IncrementBy) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increment not implemented")
}
func (UnimplementedIncrementCounterServer) mustEmbedUnimplementedIncrementCounterServer() {}

// UnsafeIncrementCounterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IncrementCounterServer will
// result in compilation errors.
type UnsafeIncrementCounterServer interface {
	mustEmbedUnimplementedIncrementCounterServer()
}

func RegisterIncrementCounterServer(s grpc.ServiceRegistrar, srv IncrementCounterServer) {
	s.RegisterService(&IncrementCounter_ServiceDesc, srv)
}

func _IncrementCounter_Increment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementBy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementCounterServer).Increment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/counter.incrementCounter/increment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementCounterServer).Increment(ctx, req.(*IncrementBy))
	}
	return interceptor(ctx, in, info, handler)
}

// IncrementCounter_ServiceDesc is the grpc.ServiceDesc for IncrementCounter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IncrementCounter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "counter.incrementCounter",
	HandlerType: (*IncrementCounterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "increment",
			Handler:    _IncrementCounter_Increment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "counter.proto",
}
