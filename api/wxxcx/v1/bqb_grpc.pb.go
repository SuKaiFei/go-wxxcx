// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: wxxcx/v1/bqb.proto

package v1

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

// BqbClient is the client API for Bqb service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BqbClient interface {
	GetBqbIndex(ctx context.Context, in *GetBqbIndexRequest, opts ...grpc.CallOption) (*GetBqbIndexReply, error)
	GetBqbList(ctx context.Context, in *GetBqbListRequest, opts ...grpc.CallOption) (*GetBqbListReply, error)
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type bqbClient struct {
	cc grpc.ClientConnInterface
}

func NewBqbClient(cc grpc.ClientConnInterface) BqbClient {
	return &bqbClient{cc}
}

func (c *bqbClient) GetBqbIndex(ctx context.Context, in *GetBqbIndexRequest, opts ...grpc.CallOption) (*GetBqbIndexReply, error) {
	out := new(GetBqbIndexReply)
	err := c.cc.Invoke(ctx, "/wxxcx.v1.Bqb/GetBqbIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bqbClient) GetBqbList(ctx context.Context, in *GetBqbListRequest, opts ...grpc.CallOption) (*GetBqbListReply, error) {
	out := new(GetBqbListReply)
	err := c.cc.Invoke(ctx, "/wxxcx.v1.Bqb/GetBqbList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bqbClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/wxxcx.v1.Bqb/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BqbServer is the server API for Bqb service.
// All implementations must embed UnimplementedBqbServer
// for forward compatibility
type BqbServer interface {
	GetBqbIndex(context.Context, *GetBqbIndexRequest) (*GetBqbIndexReply, error)
	GetBqbList(context.Context, *GetBqbListRequest) (*GetBqbListReply, error)
	Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedBqbServer()
}

// UnimplementedBqbServer must be embedded to have forward compatible implementations.
type UnimplementedBqbServer struct {
}

func (UnimplementedBqbServer) GetBqbIndex(context.Context, *GetBqbIndexRequest) (*GetBqbIndexReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBqbIndex not implemented")
}
func (UnimplementedBqbServer) GetBqbList(context.Context, *GetBqbListRequest) (*GetBqbListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBqbList not implemented")
}
func (UnimplementedBqbServer) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedBqbServer) mustEmbedUnimplementedBqbServer() {}

// UnsafeBqbServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BqbServer will
// result in compilation errors.
type UnsafeBqbServer interface {
	mustEmbedUnimplementedBqbServer()
}

func RegisterBqbServer(s grpc.ServiceRegistrar, srv BqbServer) {
	s.RegisterService(&Bqb_ServiceDesc, srv)
}

func _Bqb_GetBqbIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBqbIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BqbServer).GetBqbIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wxxcx.v1.Bqb/GetBqbIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BqbServer).GetBqbIndex(ctx, req.(*GetBqbIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bqb_GetBqbList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBqbListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BqbServer).GetBqbList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wxxcx.v1.Bqb/GetBqbList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BqbServer).GetBqbList(ctx, req.(*GetBqbListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bqb_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BqbServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wxxcx.v1.Bqb/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BqbServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Bqb_ServiceDesc is the grpc.ServiceDesc for Bqb service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bqb_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wxxcx.v1.Bqb",
	HandlerType: (*BqbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBqbIndex",
			Handler:    _Bqb_GetBqbIndex_Handler,
		},
		{
			MethodName: "GetBqbList",
			Handler:    _Bqb_GetBqbList_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Bqb_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wxxcx/v1/bqb.proto",
}
