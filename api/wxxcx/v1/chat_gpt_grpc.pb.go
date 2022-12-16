// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: wxxcx/v1/chat_gpt.proto

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

// ChatGptClient is the client API for ChatGpt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatGptClient interface {
	GetChatGptCompletions(ctx context.Context, in *GetChatGptCompletionsRequest, opts ...grpc.CallOption) (*GetChatGptCompletionsReply, error)
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type chatGptClient struct {
	cc grpc.ClientConnInterface
}

func NewChatGptClient(cc grpc.ClientConnInterface) ChatGptClient {
	return &chatGptClient{cc}
}

func (c *chatGptClient) GetChatGptCompletions(ctx context.Context, in *GetChatGptCompletionsRequest, opts ...grpc.CallOption) (*GetChatGptCompletionsReply, error) {
	out := new(GetChatGptCompletionsReply)
	err := c.cc.Invoke(ctx, "/api.wxxcx.v1.chatGpt.ChatGpt/GetChatGptCompletions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatGptClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.wxxcx.v1.chatGpt.ChatGpt/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatGptServer is the server API for ChatGpt service.
// All implementations must embed UnimplementedChatGptServer
// for forward compatibility
type ChatGptServer interface {
	GetChatGptCompletions(context.Context, *GetChatGptCompletionsRequest) (*GetChatGptCompletionsReply, error)
	Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedChatGptServer()
}

// UnimplementedChatGptServer must be embedded to have forward compatible implementations.
type UnimplementedChatGptServer struct {
}

func (UnimplementedChatGptServer) GetChatGptCompletions(context.Context, *GetChatGptCompletionsRequest) (*GetChatGptCompletionsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatGptCompletions not implemented")
}
func (UnimplementedChatGptServer) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedChatGptServer) mustEmbedUnimplementedChatGptServer() {}

// UnsafeChatGptServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatGptServer will
// result in compilation errors.
type UnsafeChatGptServer interface {
	mustEmbedUnimplementedChatGptServer()
}

func RegisterChatGptServer(s grpc.ServiceRegistrar, srv ChatGptServer) {
	s.RegisterService(&ChatGpt_ServiceDesc, srv)
}

func _ChatGpt_GetChatGptCompletions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatGptCompletionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatGptServer).GetChatGptCompletions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.wxxcx.v1.chatGpt.ChatGpt/GetChatGptCompletions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatGptServer).GetChatGptCompletions(ctx, req.(*GetChatGptCompletionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatGpt_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatGptServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.wxxcx.v1.chatGpt.ChatGpt/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatGptServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatGpt_ServiceDesc is the grpc.ServiceDesc for ChatGpt service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatGpt_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.wxxcx.v1.chatGpt.ChatGpt",
	HandlerType: (*ChatGptServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChatGptCompletions",
			Handler:    _ChatGpt_GetChatGptCompletions_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _ChatGpt_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wxxcx/v1/chat_gpt.proto",
}
