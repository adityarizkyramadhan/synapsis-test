// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: proto/category_book.proto

package grpc

import (
	context "context"
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
	CategoryBookHandler_Add_FullMethodName    = "/grpc.CategoryBookHandler/Add"
	CategoryBookHandler_Delete_FullMethodName = "/grpc.CategoryBookHandler/Delete"
)

// CategoryBookHandlerClient is the client API for CategoryBookHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryBookHandlerClient interface {
	Add(ctx context.Context, in *AddCategoryBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *DeleteCategoryBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type categoryBookHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryBookHandlerClient(cc grpc.ClientConnInterface) CategoryBookHandlerClient {
	return &categoryBookHandlerClient{cc}
}

func (c *categoryBookHandlerClient) Add(ctx context.Context, in *AddCategoryBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CategoryBookHandler_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryBookHandlerClient) Delete(ctx context.Context, in *DeleteCategoryBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CategoryBookHandler_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryBookHandlerServer is the server API for CategoryBookHandler service.
// All implementations must embed UnimplementedCategoryBookHandlerServer
// for forward compatibility.
type CategoryBookHandlerServer interface {
	Add(context.Context, *AddCategoryBookRequest) (*emptypb.Empty, error)
	Delete(context.Context, *DeleteCategoryBookRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCategoryBookHandlerServer()
}

// UnimplementedCategoryBookHandlerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCategoryBookHandlerServer struct{}

func (UnimplementedCategoryBookHandlerServer) Add(context.Context, *AddCategoryBookRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCategoryBookHandlerServer) Delete(context.Context, *DeleteCategoryBookRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCategoryBookHandlerServer) mustEmbedUnimplementedCategoryBookHandlerServer() {}
func (UnimplementedCategoryBookHandlerServer) testEmbeddedByValue()                             {}

// UnsafeCategoryBookHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryBookHandlerServer will
// result in compilation errors.
type UnsafeCategoryBookHandlerServer interface {
	mustEmbedUnimplementedCategoryBookHandlerServer()
}

func RegisterCategoryBookHandlerServer(s grpc.ServiceRegistrar, srv CategoryBookHandlerServer) {
	// If the following call pancis, it indicates UnimplementedCategoryBookHandlerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CategoryBookHandler_ServiceDesc, srv)
}

func _CategoryBookHandler_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCategoryBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryBookHandlerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryBookHandler_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryBookHandlerServer).Add(ctx, req.(*AddCategoryBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryBookHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryBookHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryBookHandler_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryBookHandlerServer).Delete(ctx, req.(*DeleteCategoryBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryBookHandler_ServiceDesc is the grpc.ServiceDesc for CategoryBookHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryBookHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.CategoryBookHandler",
	HandlerType: (*CategoryBookHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CategoryBookHandler_Add_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CategoryBookHandler_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/category_book.proto",
}