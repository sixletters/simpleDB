// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: interface/client.proto

package proto

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

// SimpleDBClient is the client API for SimpleDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpleDBClient interface {
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error)
	Put(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*KeyValue, error)
}

type simpleDBClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleDBClient(cc grpc.ClientConnInterface) SimpleDBClient {
	return &simpleDBClient{cc}
}

func (c *simpleDBClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, "/kv_store.SimpleDB/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleDBClient) Put(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*KeyValue, error) {
	out := new(KeyValue)
	err := c.cc.Invoke(ctx, "/kv_store.SimpleDB/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleDBServer is the server API for SimpleDB service.
// All implementations must embed UnimplementedSimpleDBServer
// for forward compatibility
type SimpleDBServer interface {
	Get(context.Context, *Key) (*Value, error)
	Put(context.Context, *KeyValue) (*KeyValue, error)
	mustEmbedUnimplementedSimpleDBServer()
}

// UnimplementedSimpleDBServer must be embedded to have forward compatible implementations.
type UnimplementedSimpleDBServer struct {
}

func (UnimplementedSimpleDBServer) Get(context.Context, *Key) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSimpleDBServer) Put(context.Context, *KeyValue) (*KeyValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedSimpleDBServer) mustEmbedUnimplementedSimpleDBServer() {}

// UnsafeSimpleDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpleDBServer will
// result in compilation errors.
type UnsafeSimpleDBServer interface {
	mustEmbedUnimplementedSimpleDBServer()
}

func RegisterSimpleDBServer(s grpc.ServiceRegistrar, srv SimpleDBServer) {
	s.RegisterService(&SimpleDB_ServiceDesc, srv)
}

func _SimpleDB_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleDBServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv_store.SimpleDB/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleDBServer).Get(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleDB_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleDBServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv_store.SimpleDB/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleDBServer).Put(ctx, req.(*KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

// SimpleDB_ServiceDesc is the grpc.ServiceDesc for SimpleDB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimpleDB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kv_store.SimpleDB",
	HandlerType: (*SimpleDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SimpleDB_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _SimpleDB_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interface/client.proto",
}
