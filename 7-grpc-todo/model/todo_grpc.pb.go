// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.0
// source: model/todo.proto

package model

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

// TodosClient is the client API for Todos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodosClient interface {
	CreateTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTodos(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListTodo, error)
	GetByID(ctx context.Context, in *InputTodoID, opts ...grpc.CallOption) (*Todo, error)
	UpdateTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error)
	DeleteTodo(ctx context.Context, in *InputTodoID, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type todosClient struct {
	cc grpc.ClientConnInterface
}

func NewTodosClient(cc grpc.ClientConnInterface) TodosClient {
	return &todosClient{cc}
}

func (c *todosClient) CreateTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/model.Todos/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) GetTodos(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListTodo, error) {
	out := new(ListTodo)
	err := c.cc.Invoke(ctx, "/model.Todos/GetTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) GetByID(ctx context.Context, in *InputTodoID, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/model.Todos/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) UpdateTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/model.Todos/UpdateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) DeleteTodo(ctx context.Context, in *InputTodoID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/model.Todos/DeleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodosServer is the server API for Todos service.
// All implementations must embed UnimplementedTodosServer
// for forward compatibility
type TodosServer interface {
	CreateTodo(context.Context, *Todo) (*emptypb.Empty, error)
	GetTodos(context.Context, *emptypb.Empty) (*ListTodo, error)
	GetByID(context.Context, *InputTodoID) (*Todo, error)
	UpdateTodo(context.Context, *Todo) (*Todo, error)
	DeleteTodo(context.Context, *InputTodoID) (*emptypb.Empty, error)
	mustEmbedUnimplementedTodosServer()
}

// UnimplementedTodosServer must be embedded to have forward compatible implementations.
type UnimplementedTodosServer struct {
}

func (UnimplementedTodosServer) CreateTodo(context.Context, *Todo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedTodosServer) GetTodos(context.Context, *emptypb.Empty) (*ListTodo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodos not implemented")
}
func (UnimplementedTodosServer) GetByID(context.Context, *InputTodoID) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedTodosServer) UpdateTodo(context.Context, *Todo) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (UnimplementedTodosServer) DeleteTodo(context.Context, *InputTodoID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}
func (UnimplementedTodosServer) mustEmbedUnimplementedTodosServer() {}

// UnsafeTodosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodosServer will
// result in compilation errors.
type UnsafeTodosServer interface {
	mustEmbedUnimplementedTodosServer()
}

func RegisterTodosServer(s grpc.ServiceRegistrar, srv TodosServer) {
	s.RegisterService(&Todos_ServiceDesc, srv)
}

func _Todos_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).CreateTodo(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_GetTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).GetTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/GetTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).GetTodos(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputTodoID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).GetByID(ctx, req.(*InputTodoID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/UpdateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).UpdateTodo(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputTodoID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).DeleteTodo(ctx, req.(*InputTodoID))
	}
	return interceptor(ctx, in, info, handler)
}

// Todos_ServiceDesc is the grpc.ServiceDesc for Todos service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Todos_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.Todos",
	HandlerType: (*TodosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _Todos_CreateTodo_Handler,
		},
		{
			MethodName: "GetTodos",
			Handler:    _Todos_GetTodos_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _Todos_GetByID_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _Todos_UpdateTodo_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _Todos_DeleteTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model/todo.proto",
}
