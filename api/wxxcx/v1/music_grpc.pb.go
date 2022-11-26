// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: wxxcx/v1/music.proto

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

// MusicClient is the client API for Music service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MusicClient interface {
	GetMusicList(ctx context.Context, in *GetMusicListRequest, opts ...grpc.CallOption) (*GetMusicListReply, error)
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type musicClient struct {
	cc grpc.ClientConnInterface
}

func NewMusicClient(cc grpc.ClientConnInterface) MusicClient {
	return &musicClient{cc}
}

func (c *musicClient) GetMusicList(ctx context.Context, in *GetMusicListRequest, opts ...grpc.CallOption) (*GetMusicListReply, error) {
	out := new(GetMusicListReply)
	err := c.cc.Invoke(ctx, "/wxxcx.v1.music.Music/GetMusicList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/wxxcx.v1.music.Music/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MusicServer is the server API for Music service.
// All implementations must embed UnimplementedMusicServer
// for forward compatibility
type MusicServer interface {
	GetMusicList(context.Context, *GetMusicListRequest) (*GetMusicListReply, error)
	Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedMusicServer()
}

// UnimplementedMusicServer must be embedded to have forward compatible implementations.
type UnimplementedMusicServer struct {
}

func (UnimplementedMusicServer) GetMusicList(context.Context, *GetMusicListRequest) (*GetMusicListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMusicList not implemented")
}
func (UnimplementedMusicServer) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedMusicServer) mustEmbedUnimplementedMusicServer() {}

// UnsafeMusicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MusicServer will
// result in compilation errors.
type UnsafeMusicServer interface {
	mustEmbedUnimplementedMusicServer()
}

func RegisterMusicServer(s grpc.ServiceRegistrar, srv MusicServer) {
	s.RegisterService(&Music_ServiceDesc, srv)
}

func _Music_GetMusicList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMusicListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServer).GetMusicList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wxxcx.v1.music.Music/GetMusicList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServer).GetMusicList(ctx, req.(*GetMusicListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Music_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wxxcx.v1.music.Music/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Music_ServiceDesc is the grpc.ServiceDesc for Music service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Music_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wxxcx.v1.music.Music",
	HandlerType: (*MusicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMusicList",
			Handler:    _Music_GetMusicList_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Music_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wxxcx/v1/music.proto",
}
