// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: service/api.service.proto

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	TaskCreate(ctx context.Context, in *TaskCreateRequest, opts ...grpc.CallOption) (*TaskCreateResponse, error)
	TaskList(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (*TaskListResponse, error)
	TaskUpdate(ctx context.Context, in *TaskUpdateRequest, opts ...grpc.CallOption) (*TaskUpdateResponse, error)
	TaskDelete(ctx context.Context, in *TaskDeleteRequest, opts ...grpc.CallOption) (*TaskDeleteResponse, error)
	TaskGet(ctx context.Context, in *TaskGetRequest, opts ...grpc.CallOption) (*TaskGetResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) TaskCreate(ctx context.Context, in *TaskCreateRequest, opts ...grpc.CallOption) (*TaskCreateResponse, error) {
	out := new(TaskCreateResponse)
	err := c.cc.Invoke(ctx, "/ozon.dev.vanek623.task_manager_bot.api.service.Service/TaskCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) TaskList(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (*TaskListResponse, error) {
	out := new(TaskListResponse)
	err := c.cc.Invoke(ctx, "/ozon.dev.vanek623.task_manager_bot.api.service.Service/TaskList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) TaskUpdate(ctx context.Context, in *TaskUpdateRequest, opts ...grpc.CallOption) (*TaskUpdateResponse, error) {
	out := new(TaskUpdateResponse)
	err := c.cc.Invoke(ctx, "/ozon.dev.vanek623.task_manager_bot.api.service.Service/TaskUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) TaskDelete(ctx context.Context, in *TaskDeleteRequest, opts ...grpc.CallOption) (*TaskDeleteResponse, error) {
	out := new(TaskDeleteResponse)
	err := c.cc.Invoke(ctx, "/ozon.dev.vanek623.task_manager_bot.api.service.Service/TaskDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) TaskGet(ctx context.Context, in *TaskGetRequest, opts ...grpc.CallOption) (*TaskGetResponse, error) {
	out := new(TaskGetResponse)
	err := c.cc.Invoke(ctx, "/ozon.dev.vanek623.task_manager_bot.api.service.Service/TaskGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	TaskCreate(context.Context, *TaskCreateRequest) (*TaskCreateResponse, error)
	TaskList(context.Context, *TaskListRequest) (*TaskListResponse, error)
	TaskUpdate(context.Context, *TaskUpdateRequest) (*TaskUpdateResponse, error)
	TaskDelete(context.Context, *TaskDeleteRequest) (*TaskDeleteResponse, error)
	TaskGet(context.Context, *TaskGetRequest) (*TaskGetResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) TaskCreate(context.Context, *TaskCreateRequest) (*TaskCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskCreate not implemented")
}
func (UnimplementedServiceServer) TaskList(context.Context, *TaskListRequest) (*TaskListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskList not implemented")
}
func (UnimplementedServiceServer) TaskUpdate(context.Context, *TaskUpdateRequest) (*TaskUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskUpdate not implemented")
}
func (UnimplementedServiceServer) TaskDelete(context.Context, *TaskDeleteRequest) (*TaskDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskDelete not implemented")
}
func (UnimplementedServiceServer) TaskGet(context.Context, *TaskGetRequest) (*TaskGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskGet not implemented")
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

func _Service_TaskCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).TaskCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.vanek623.task_manager_bot.api.service.Service/TaskCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).TaskCreate(ctx, req.(*TaskCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_TaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).TaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.vanek623.task_manager_bot.api.service.Service/TaskList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).TaskList(ctx, req.(*TaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_TaskUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).TaskUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.vanek623.task_manager_bot.api.service.Service/TaskUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).TaskUpdate(ctx, req.(*TaskUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_TaskDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).TaskDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.vanek623.task_manager_bot.api.service.Service/TaskDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).TaskDelete(ctx, req.(*TaskDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_TaskGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).TaskGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.vanek623.task_manager_bot.api.service.Service/TaskGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).TaskGet(ctx, req.(*TaskGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ozon.dev.vanek623.task_manager_bot.api.service.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TaskCreate",
			Handler:    _Service_TaskCreate_Handler,
		},
		{
			MethodName: "TaskList",
			Handler:    _Service_TaskList_Handler,
		},
		{
			MethodName: "TaskUpdate",
			Handler:    _Service_TaskUpdate_Handler,
		},
		{
			MethodName: "TaskDelete",
			Handler:    _Service_TaskDelete_Handler,
		},
		{
			MethodName: "TaskGet",
			Handler:    _Service_TaskGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/api.service.proto",
}
