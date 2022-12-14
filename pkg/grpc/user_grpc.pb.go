// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: user.proto

package grpc

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	User(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	UserServerStream(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (UserService_UserServerStreamClient, error)
	UserClientStream(ctx context.Context, opts ...grpc.CallOption) (UserService_UserClientStreamClient, error)
	UserBidirectStream(ctx context.Context, opts ...grpc.CallOption) (UserService_UserBidirectStreamClient, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) User(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/grpcapp.UserService/User", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserServerStream(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (UserService_UserServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], "/grpcapp.UserService/UserServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceUserServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_UserServerStreamClient interface {
	Recv() (*UserResponse, error)
	grpc.ClientStream
}

type userServiceUserServerStreamClient struct {
	grpc.ClientStream
}

func (x *userServiceUserServerStreamClient) Recv() (*UserResponse, error) {
	m := new(UserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) UserClientStream(ctx context.Context, opts ...grpc.CallOption) (UserService_UserClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[1], "/grpcapp.UserService/UserClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceUserClientStreamClient{stream}
	return x, nil
}

type UserService_UserClientStreamClient interface {
	Send(*UserRequest) error
	CloseAndRecv() (*UserResponse, error)
	grpc.ClientStream
}

type userServiceUserClientStreamClient struct {
	grpc.ClientStream
}

func (x *userServiceUserClientStreamClient) Send(m *UserRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceUserClientStreamClient) CloseAndRecv() (*UserResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) UserBidirectStream(ctx context.Context, opts ...grpc.CallOption) (UserService_UserBidirectStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[2], "/grpcapp.UserService/UserBidirectStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceUserBidirectStreamClient{stream}
	return x, nil
}

type UserService_UserBidirectStreamClient interface {
	Send(*UserRequest) error
	Recv() (*UserResponse, error)
	grpc.ClientStream
}

type userServiceUserBidirectStreamClient struct {
	grpc.ClientStream
}

func (x *userServiceUserBidirectStreamClient) Send(m *UserRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceUserBidirectStreamClient) Recv() (*UserResponse, error) {
	m := new(UserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	User(context.Context, *UserRequest) (*UserResponse, error)
	UserServerStream(*UserRequest, UserService_UserServerStreamServer) error
	UserClientStream(UserService_UserClientStreamServer) error
	UserBidirectStream(UserService_UserBidirectStreamServer) error
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) User(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method User not implemented")
}
func (UnimplementedUserServiceServer) UserServerStream(*UserRequest, UserService_UserServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method UserServerStream not implemented")
}
func (UnimplementedUserServiceServer) UserClientStream(UserService_UserClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method UserClientStream not implemented")
}
func (UnimplementedUserServiceServer) UserBidirectStream(UserService_UserBidirectStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method UserBidirectStream not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).User(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapp.UserService/User",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).User(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).UserServerStream(m, &userServiceUserServerStreamServer{stream})
}

type UserService_UserServerStreamServer interface {
	Send(*UserResponse) error
	grpc.ServerStream
}

type userServiceUserServerStreamServer struct {
	grpc.ServerStream
}

func (x *userServiceUserServerStreamServer) Send(m *UserResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_UserClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).UserClientStream(&userServiceUserClientStreamServer{stream})
}

type UserService_UserClientStreamServer interface {
	SendAndClose(*UserResponse) error
	Recv() (*UserRequest, error)
	grpc.ServerStream
}

type userServiceUserClientStreamServer struct {
	grpc.ServerStream
}

func (x *userServiceUserClientStreamServer) SendAndClose(m *UserResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceUserClientStreamServer) Recv() (*UserRequest, error) {
	m := new(UserRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UserService_UserBidirectStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).UserBidirectStream(&userServiceUserBidirectStreamServer{stream})
}

type UserService_UserBidirectStreamServer interface {
	Send(*UserResponse) error
	Recv() (*UserRequest, error)
	grpc.ServerStream
}

type userServiceUserBidirectStreamServer struct {
	grpc.ServerStream
}

func (x *userServiceUserBidirectStreamServer) Send(m *UserResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceUserBidirectStreamServer) Recv() (*UserRequest, error) {
	m := new(UserRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapp.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "User",
			Handler:    _UserService_User_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UserServerStream",
			Handler:       _UserService_UserServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UserClientStream",
			Handler:       _UserService_UserClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UserBidirectStream",
			Handler:       _UserService_UserBidirectStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "user.proto",
}
