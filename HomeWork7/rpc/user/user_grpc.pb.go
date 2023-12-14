// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: user.proto

package user

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
	RServer_Register_FullMethodName = "/RServer/Register"
)

// RServerClient is the client API for RServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RServerClient interface {
	Register(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Response, error)
}

type rServerClient struct {
	cc grpc.ClientConnInterface
}

func NewRServerClient(cc grpc.ClientConnInterface) RServerClient {
	return &rServerClient{cc}
}

func (c *rServerClient) Register(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, RServer_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RServerServer is the server API for RServer service.
// All implementations must embed UnimplementedRServerServer
// for forward compatibility
type RServerServer interface {
	Register(context.Context, *UserInfo) (*Response, error)
	mustEmbedUnimplementedRServerServer()
}

// UnimplementedRServerServer must be embedded to have forward compatible implementations.
type UnimplementedRServerServer struct {
}

func (UnimplementedRServerServer) Register(context.Context, *UserInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRServerServer) mustEmbedUnimplementedRServerServer() {}

// UnsafeRServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RServerServer will
// result in compilation errors.
type UnsafeRServerServer interface {
	mustEmbedUnimplementedRServerServer()
}

func RegisterRServerServer(s grpc.ServiceRegistrar, srv RServerServer) {
	s.RegisterService(&RServer_ServiceDesc, srv)
}

func _RServer_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RServerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RServer_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RServerServer).Register(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// RServer_ServiceDesc is the grpc.ServiceDesc for RServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RServer",
	HandlerType: (*RServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _RServer_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

const (
	LServer_Login_FullMethodName = "/LServer/Login"
)

// LServerClient is the client API for LServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LServerClient interface {
	Login(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Response, error)
}

type lServerClient struct {
	cc grpc.ClientConnInterface
}

func NewLServerClient(cc grpc.ClientConnInterface) LServerClient {
	return &lServerClient{cc}
}

func (c *lServerClient) Login(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, LServer_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LServerServer is the server API for LServer service.
// All implementations must embed UnimplementedLServerServer
// for forward compatibility
type LServerServer interface {
	Login(context.Context, *UserInfo) (*Response, error)
	mustEmbedUnimplementedLServerServer()
}

// UnimplementedLServerServer must be embedded to have forward compatible implementations.
type UnimplementedLServerServer struct {
}

func (UnimplementedLServerServer) Login(context.Context, *UserInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedLServerServer) mustEmbedUnimplementedLServerServer() {}

// UnsafeLServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LServerServer will
// result in compilation errors.
type UnsafeLServerServer interface {
	mustEmbedUnimplementedLServerServer()
}

func RegisterLServerServer(s grpc.ServiceRegistrar, srv LServerServer) {
	s.RegisterService(&LServer_ServiceDesc, srv)
}

func _LServer_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LServerServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LServer_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LServerServer).Login(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// LServer_ServiceDesc is the grpc.ServiceDesc for LServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LServer",
	HandlerType: (*LServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _LServer_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}