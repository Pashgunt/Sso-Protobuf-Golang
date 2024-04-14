// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: permission/permission.proto

package sso

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PermissionClient is the client API for Permission service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionClient interface {
	AddNewPermission(ctx context.Context, in *AddNewPermissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemovePermission(ctx context.Context, in *RemovePermissionRequest, opts ...grpc.CallOption) (*RemovePermissionResponse, error)
	UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*UpdatePermissionResponse, error)
	PermissionList(ctx context.Context, in *PermissionListRequest, opts ...grpc.CallOption) (*PermissionListResponse, error)
}

type permissionClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionClient(cc grpc.ClientConnInterface) PermissionClient {
	return &permissionClient{cc}
}

func (c *permissionClient) AddNewPermission(ctx context.Context, in *AddNewPermissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/permission.Permission/AddNewPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionClient) RemovePermission(ctx context.Context, in *RemovePermissionRequest, opts ...grpc.CallOption) (*RemovePermissionResponse, error) {
	out := new(RemovePermissionResponse)
	err := c.cc.Invoke(ctx, "/permission.Permission/RemovePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionClient) UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*UpdatePermissionResponse, error) {
	out := new(UpdatePermissionResponse)
	err := c.cc.Invoke(ctx, "/permission.Permission/UpdatePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionClient) PermissionList(ctx context.Context, in *PermissionListRequest, opts ...grpc.CallOption) (*PermissionListResponse, error) {
	out := new(PermissionListResponse)
	err := c.cc.Invoke(ctx, "/permission.Permission/PermissionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionServer is the server API for Permission service.
// All implementations must embed UnimplementedPermissionServer
// for forward compatibility
type PermissionServer interface {
	AddNewPermission(context.Context, *AddNewPermissionRequest) (*emptypb.Empty, error)
	RemovePermission(context.Context, *RemovePermissionRequest) (*RemovePermissionResponse, error)
	UpdatePermission(context.Context, *UpdatePermissionRequest) (*UpdatePermissionResponse, error)
	PermissionList(context.Context, *PermissionListRequest) (*PermissionListResponse, error)
	mustEmbedUnimplementedPermissionServer()
}

// UnimplementedPermissionServer must be embedded to have forward compatible implementations.
type UnimplementedPermissionServer struct {
}

func (UnimplementedPermissionServer) AddNewPermission(context.Context, *AddNewPermissionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewPermission not implemented")
}
func (UnimplementedPermissionServer) RemovePermission(context.Context, *RemovePermissionRequest) (*RemovePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePermission not implemented")
}
func (UnimplementedPermissionServer) UpdatePermission(context.Context, *UpdatePermissionRequest) (*UpdatePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePermission not implemented")
}
func (UnimplementedPermissionServer) PermissionList(context.Context, *PermissionListRequest) (*PermissionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionList not implemented")
}
func (UnimplementedPermissionServer) mustEmbedUnimplementedPermissionServer() {}

// UnsafePermissionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionServer will
// result in compilation errors.
type UnsafePermissionServer interface {
	mustEmbedUnimplementedPermissionServer()
}

func RegisterPermissionServer(s grpc.ServiceRegistrar, srv PermissionServer) {
	s.RegisterService(&Permission_ServiceDesc, srv)
}

func _Permission_AddNewPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNewPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServer).AddNewPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permission.Permission/AddNewPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServer).AddNewPermission(ctx, req.(*AddNewPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Permission_RemovePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServer).RemovePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permission.Permission/RemovePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServer).RemovePermission(ctx, req.(*RemovePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Permission_UpdatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServer).UpdatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permission.Permission/UpdatePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServer).UpdatePermission(ctx, req.(*UpdatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Permission_PermissionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServer).PermissionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permission.Permission/PermissionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServer).PermissionList(ctx, req.(*PermissionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Permission_ServiceDesc is the grpc.ServiceDesc for Permission service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Permission_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "permission.Permission",
	HandlerType: (*PermissionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNewPermission",
			Handler:    _Permission_AddNewPermission_Handler,
		},
		{
			MethodName: "RemovePermission",
			Handler:    _Permission_RemovePermission_Handler,
		},
		{
			MethodName: "UpdatePermission",
			Handler:    _Permission_UpdatePermission_Handler,
		},
		{
			MethodName: "PermissionList",
			Handler:    _Permission_PermissionList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "permission/permission.proto",
}