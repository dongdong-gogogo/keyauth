// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: apps/token/pb/token.proto

package token

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// 颁发Token(Login)
	IssueToken(ctx context.Context, in *IssueTokenRequest, opts ...grpc.CallOption) (*Token, error)
	// 撤销Token(Logout)
	RevolkToken(ctx context.Context, in *RevolkTokenRequest, opts ...grpc.CallOption) (*Token, error)
	// 校验Token的接口(内部服务使用)
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*Token, error)
	// 查询Token, 查询用于REST ful API访问颁发出去的Token
	QueryToken(ctx context.Context, in *QueryTokenRequest, opts ...grpc.CallOption) (*TokenSet, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) IssueToken(ctx context.Context, in *IssueTokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/keyauth.token.Service/IssueToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) RevolkToken(ctx context.Context, in *RevolkTokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/keyauth.token.Service/RevolkToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/keyauth.token.Service/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryToken(ctx context.Context, in *QueryTokenRequest, opts ...grpc.CallOption) (*TokenSet, error) {
	out := new(TokenSet)
	err := c.cc.Invoke(ctx, "/keyauth.token.Service/QueryToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// 颁发Token(Login)
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	// 校验Token的接口(内部服务使用)
	// 撤销Token(Logout)
	RevolkToken(context.Context, *RevolkTokenRequest) (*Token, error)
	ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error)
	// 查询Token, 查询用于REST ful API访问颁发出去的Token
	QueryToken(context.Context, *QueryTokenRequest) (*TokenSet, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) IssueToken(context.Context, *IssueTokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueToken not implemented")
}
func (UnimplementedServiceServer) RevolkToken(context.Context, *RevolkTokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevolkToken not implemented")
}
func (UnimplementedServiceServer) ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedServiceServer) QueryToken(context.Context, *QueryTokenRequest) (*TokenSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryToken not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_IssueToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).IssueToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.token.Service/IssueToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).IssueToken(ctx, req.(*IssueTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_RevolkToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevolkTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).RevolkToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.token.Service/RevolkToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).RevolkToken(ctx, req.(*RevolkTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.token.Service/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ValidateToken(ctx, req.(*ValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.token.Service/QueryToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryToken(ctx, req.(*QueryTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keyauth.token.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueToken",
			Handler:    _Service_IssueToken_Handler,
		},
		{
			MethodName: "RevolkToken",
			Handler:    _Service_RevolkToken_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _Service_ValidateToken_Handler,
		},
		{
			MethodName: "QueryToken",
			Handler:    _Service_QueryToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/token/pb/token.proto",
}
