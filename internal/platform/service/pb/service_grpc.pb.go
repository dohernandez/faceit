// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: service.proto

package api

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
	FaceitService_AddUser_FullMethodName    = "/api.faceit.FaceitService/AddUser"
	FaceitService_UpdateUser_FullMethodName = "/api.faceit.FaceitService/UpdateUser"
)

// FaceitServiceClient is the client API for FaceitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The API to manages users.
type FaceitServiceClient interface {
	// AddUser add new user.
	//
	// Receives a request with user data. Responses whether the user was added successfully or not.
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
	// Update the user.
	//
	// Receives a request with user data. Responses whether the user was updated successfully or not.
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type faceitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFaceitServiceClient(cc grpc.ClientConnInterface) FaceitServiceClient {
	return &faceitServiceClient{cc}
}

func (c *faceitServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddUserResponse)
	err := c.cc.Invoke(ctx, FaceitService_AddUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faceitServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FaceitService_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FaceitServiceServer is the server API for FaceitService service.
// All implementations must embed UnimplementedFaceitServiceServer
// for forward compatibility.
//
// The API to manages users.
type FaceitServiceServer interface {
	// AddUser add new user.
	//
	// Receives a request with user data. Responses whether the user was added successfully or not.
	AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error)
	// Update the user.
	//
	// Receives a request with user data. Responses whether the user was updated successfully or not.
	UpdateUser(context.Context, *UpdateUserRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedFaceitServiceServer()
}

// UnimplementedFaceitServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFaceitServiceServer struct{}

func (UnimplementedFaceitServiceServer) AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedFaceitServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedFaceitServiceServer) mustEmbedUnimplementedFaceitServiceServer() {}
func (UnimplementedFaceitServiceServer) testEmbeddedByValue()                       {}

// UnsafeFaceitServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FaceitServiceServer will
// result in compilation errors.
type UnsafeFaceitServiceServer interface {
	mustEmbedUnimplementedFaceitServiceServer()
}

func RegisterFaceitServiceServer(s grpc.ServiceRegistrar, srv FaceitServiceServer) {
	// If the following call pancis, it indicates UnimplementedFaceitServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FaceitService_ServiceDesc, srv)
}

func _FaceitService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaceitServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FaceitService_AddUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaceitServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FaceitService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaceitServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FaceitService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaceitServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FaceitService_ServiceDesc is the grpc.ServiceDesc for FaceitService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FaceitService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.faceit.FaceitService",
	HandlerType: (*FaceitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _FaceitService_AddUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _FaceitService_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
