// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: service_simple_wallet.proto

package pb

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
	SimpleWallet_CreateUser_FullMethodName = "/pb.SimpleWallet/createUser"
	SimpleWallet_UpdateUser_FullMethodName = "/pb.SimpleWallet/updateUser"
	SimpleWallet_Login_FullMethodName      = "/pb.SimpleWallet/login"
)

// SimpleWalletClient is the client API for SimpleWallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpleWalletClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	Login(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type simpleWalletClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleWalletClient(cc grpc.ClientConnInterface) SimpleWalletClient {
	return &simpleWalletClient{cc}
}

func (c *simpleWalletClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, SimpleWallet_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleWalletClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, SimpleWallet_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleWalletClient) Login(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, SimpleWallet_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleWalletServer is the server API for SimpleWallet service.
// All implementations must embed UnimplementedSimpleWalletServer
// for forward compatibility
type SimpleWalletServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	Login(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	mustEmbedUnimplementedSimpleWalletServer()
}

// UnimplementedSimpleWalletServer must be embedded to have forward compatible implementations.
type UnimplementedSimpleWalletServer struct {
}

func (UnimplementedSimpleWalletServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedSimpleWalletServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedSimpleWalletServer) Login(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSimpleWalletServer) mustEmbedUnimplementedSimpleWalletServer() {}

// UnsafeSimpleWalletServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpleWalletServer will
// result in compilation errors.
type UnsafeSimpleWalletServer interface {
	mustEmbedUnimplementedSimpleWalletServer()
}

func RegisterSimpleWalletServer(s grpc.ServiceRegistrar, srv SimpleWalletServer) {
	s.RegisterService(&SimpleWallet_ServiceDesc, srv)
}

func _SimpleWallet_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleWalletServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleWallet_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleWalletServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleWallet_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleWalletServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleWallet_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleWalletServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleWallet_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleWalletServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleWallet_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleWalletServer).Login(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SimpleWallet_ServiceDesc is the grpc.ServiceDesc for SimpleWallet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimpleWallet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SimpleWallet",
	HandlerType: (*SimpleWalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createUser",
			Handler:    _SimpleWallet_CreateUser_Handler,
		},
		{
			MethodName: "updateUser",
			Handler:    _SimpleWallet_UpdateUser_Handler,
		},
		{
			MethodName: "login",
			Handler:    _SimpleWallet_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_simple_wallet.proto",
}
