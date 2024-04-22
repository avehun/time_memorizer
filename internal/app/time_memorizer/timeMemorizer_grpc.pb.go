// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

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

// TimeMemorizerClient is the client API for TimeMemorizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeMemorizerClient interface {
	AddTime(ctx context.Context, in *CategoryAndTime, opts ...grpc.CallOption) (*Message, error)
	SubstractTime(ctx context.Context, in *CategoryAndTime, opts ...grpc.CallOption) (*Message, error)
}

type timeMemorizerClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeMemorizerClient(cc grpc.ClientConnInterface) TimeMemorizerClient {
	return &timeMemorizerClient{cc}
}

func (c *timeMemorizerClient) AddTime(ctx context.Context, in *CategoryAndTime, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/TimeMemorizer/AddTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeMemorizerClient) SubstractTime(ctx context.Context, in *CategoryAndTime, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/TimeMemorizer/SubstractTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeMemorizerServer is the server API for TimeMemorizer service.
// All implementations must embed UnimplementedTimeMemorizerServer
// for forward compatibility
type TimeMemorizerServer interface {
	AddTime(context.Context, *CategoryAndTime) (*Message, error)
	SubstractTime(context.Context, *CategoryAndTime) (*Message, error)
	mustEmbedUnimplementedTimeMemorizerServer()
}

// UnimplementedTimeMemorizerServer must be embedded to have forward compatible implementations.
type UnimplementedTimeMemorizerServer struct {
}

func (UnimplementedTimeMemorizerServer) AddTime(context.Context, *CategoryAndTime) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTime not implemented")
}
func (UnimplementedTimeMemorizerServer) SubstractTime(context.Context, *CategoryAndTime) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubstractTime not implemented")
}
func (UnimplementedTimeMemorizerServer) mustEmbedUnimplementedTimeMemorizerServer() {}

// UnsafeTimeMemorizerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeMemorizerServer will
// result in compilation errors.
type UnsafeTimeMemorizerServer interface {
	mustEmbedUnimplementedTimeMemorizerServer()
}

func RegisterTimeMemorizerServer(s grpc.ServiceRegistrar, srv TimeMemorizerServer) {
	s.RegisterService(&TimeMemorizer_ServiceDesc, srv)
}

func _TimeMemorizer_AddTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryAndTime)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeMemorizerServer).AddTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TimeMemorizer/AddTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeMemorizerServer).AddTime(ctx, req.(*CategoryAndTime))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeMemorizer_SubstractTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryAndTime)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeMemorizerServer).SubstractTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TimeMemorizer/SubstractTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeMemorizerServer).SubstractTime(ctx, req.(*CategoryAndTime))
	}
	return interceptor(ctx, in, info, handler)
}

// TimeMemorizer_ServiceDesc is the grpc.ServiceDesc for TimeMemorizer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeMemorizer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TimeMemorizer",
	HandlerType: (*TimeMemorizerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTime",
			Handler:    _TimeMemorizer_AddTime_Handler,
		},
		{
			MethodName: "SubstractTime",
			Handler:    _TimeMemorizer_SubstractTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "timeMemorizer.proto",
}
