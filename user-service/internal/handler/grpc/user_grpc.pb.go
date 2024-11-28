// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: proto/user.proto

package user

import (
	context "context"

	"github.com/adityarizkyramadhan/synapsis-test/user-service/internal/service"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserHandler_GetByID_FullMethodName = "/user.UserHandler/GetByID"
	UserHandler_Create_FullMethodName  = "/user.UserHandler/Create"
	UserHandler_Update_FullMethodName  = "/user.UserHandler/Update"
	UserHandler_Delete_FullMethodName  = "/user.UserHandler/Delete"
)

// UserHandlerClient is the client API for UserHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserHandlerClient interface {
	GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*User, error)
	Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserHandlerClient(cc grpc.ClientConnInterface) UserHandlerClient {
	return &userHandlerClient{cc}
}

func (c *userHandlerClient) GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, UserHandler_GetByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userHandlerClient) Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserHandler_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userHandlerClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserHandler_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userHandlerClient) Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserHandler_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserHandlerServer is the server API for UserHandler service.
// All implementations must embed UnimplementedUserHandlerServer
// for forward compatibility.
type UserHandlerServer interface {
	GetByID(context.Context, *GetByIDRequest) (*User, error)
	Create(context.Context, *User) (*emptypb.Empty, error)
	Update(context.Context, *UpdateUserRequest) (*emptypb.Empty, error)
	Delete(context.Context, *DeleteUserRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserHandlerServer()
}

// UnimplementedUserHandlerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UserHandlerServerStruct struct{
	serviceUser service.UserService
}

func NewUserHandlerServer(serviceUser service.UserService) UserHandlerServer {
	return &UserHandlerServerStruct{serviceUser: serviceUser}
}

func (u *UserHandlerServerStruct) GetByID(context.Context, *GetByIDRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (u *UserHandlerServerStruct) Create(context.Context, *User) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (u *UserHandlerServerStruct) Update(context.Context, *UpdateUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (u *UserHandlerServerStruct) Delete(context.Context, *DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UserHandlerServerStruct) mustEmbedUnimplementedUserHandlerServer() {}
func (UserHandlerServerStruct) testEmbeddedByValue()                     {}

// UnsafeUserHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserHandlerServer will
// result in compilation errors.
type UnsafeUserHandlerServer interface {
	mustEmbedUnimplementedUserHandlerServer()
}

func RegisterUserHandlerServer(s grpc.ServiceRegistrar, srv UserHandlerServer) {
	// If the following call pancis, it indicates UnimplementedUserHandlerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserHandler_ServiceDesc, srv)
}

func _UserHandler_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHandlerServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserHandler_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHandlerServer).GetByID(ctx, req.(*GetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserHandler_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHandlerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserHandler_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHandlerServer).Create(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserHandler_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHandlerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserHandler_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHandlerServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserHandler_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHandlerServer).Delete(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserHandler_ServiceDesc is the grpc.ServiceDesc for UserHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserHandler",
	HandlerType: (*UserHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByID",
			Handler:    _UserHandler_GetByID_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UserHandler_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserHandler_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserHandler_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}
