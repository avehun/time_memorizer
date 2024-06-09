// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: timeMemorizer.proto

package api

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

const (
	TimeMemorizer_AddTime_FullMethodName       = "/TimeMemorizer/AddTime"
	TimeMemorizer_SubstractTime_FullMethodName = "/TimeMemorizer/SubstractTime"
	TimeMemorizer_ShowTime_FullMethodName      = "/TimeMemorizer/ShowTime"
)

// TimeMemorizerClient is the client API for TimeMemorizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeMemorizerClient interface {
	AddTime(ctx context.Context, in *AddTimeRequest, opts ...grpc.CallOption) (*AddTimeResponse, error)
	SubstractTime(ctx context.Context, in *SubtractTimeRequest, opts ...grpc.CallOption) (*SubtractTimeResponse, error)
	ShowTime(ctx context.Context, in *ShowTimeRequest, opts ...grpc.CallOption) (*ShowTimeResponse, error)
}

type timeMemorizerClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeMemorizerClient(cc grpc.ClientConnInterface) TimeMemorizerClient {
	return &timeMemorizerClient{cc}
}

func (c *timeMemorizerClient) AddTime(ctx context.Context, in *AddTimeRequest, opts ...grpc.CallOption) (*AddTimeResponse, error) {
	out := new(AddTimeResponse)
	err := c.cc.Invoke(ctx, TimeMemorizer_AddTime_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeMemorizerClient) SubstractTime(ctx context.Context, in *SubtractTimeRequest, opts ...grpc.CallOption) (*SubtractTimeResponse, error) {
	out := new(SubtractTimeResponse)
	err := c.cc.Invoke(ctx, TimeMemorizer_SubstractTime_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeMemorizerClient) ShowTime(ctx context.Context, in *ShowTimeRequest, opts ...grpc.CallOption) (*ShowTimeResponse, error) {
	out := new(ShowTimeResponse)
	err := c.cc.Invoke(ctx, TimeMemorizer_ShowTime_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeMemorizerServer is the server API for TimeMemorizer service.
// All implementations must embed UnimplementedTimeMemorizerServer
// for forward compatibility
type TimeMemorizerServer interface {
	AddTime(context.Context, *AddTimeRequest) (*AddTimeResponse, error)
	SubstractTime(context.Context, *SubtractTimeRequest) (*SubtractTimeResponse, error)
	ShowTime(context.Context, *ShowTimeRequest) (*ShowTimeResponse, error)
	mustEmbedUnimplementedTimeMemorizerServer()
}

// UnimplementedTimeMemorizerServer must be embedded to have forward compatible implementations.
type UnimplementedTimeMemorizerServer struct {
}

func (UnimplementedTimeMemorizerServer) AddTime(context.Context, *AddTimeRequest) (*AddTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTime not implemented")
}
func (UnimplementedTimeMemorizerServer) SubstractTime(context.Context, *SubtractTimeRequest) (*SubtractTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubstractTime not implemented")
}
func (UnimplementedTimeMemorizerServer) ShowTime(context.Context, *ShowTimeRequest) (*ShowTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowTime not implemented")
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
	in := new(AddTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeMemorizerServer).AddTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeMemorizer_AddTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeMemorizerServer).AddTime(ctx, req.(*AddTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeMemorizer_SubstractTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubtractTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeMemorizerServer).SubstractTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeMemorizer_SubstractTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeMemorizerServer).SubstractTime(ctx, req.(*SubtractTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeMemorizer_ShowTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeMemorizerServer).ShowTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeMemorizer_ShowTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeMemorizerServer).ShowTime(ctx, req.(*ShowTimeRequest))
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
		{
			MethodName: "ShowTime",
			Handler:    _TimeMemorizer_ShowTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "timeMemorizer.proto",
}