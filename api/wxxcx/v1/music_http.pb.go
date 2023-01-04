// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type MusicHTTPServer interface {
	GetMusicList(context.Context, *GetMusicListRequest) (*GetMusicListReply, error)
	Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
}

func RegisterMusicHTTPServer(s *http.Server, srv MusicHTTPServer) {
	r := s.Route("/")
	r.GET("/wxxcx/music/list", _Music_GetMusicList0_HTTP_Handler(srv))
	r.GET("/wxxcx/music/ping", _Music_Ping5_HTTP_Handler(srv))
}

func _Music_GetMusicList0_HTTP_Handler(srv MusicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMusicListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/wxxcx.v1.music.Music/GetMusicList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMusicList(ctx, req.(*GetMusicListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMusicListReply)
		return ctx.Result(200, reply)
	}
}

func _Music_Ping5_HTTP_Handler(srv MusicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/wxxcx.v1.music.Music/Ping")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Ping(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type MusicHTTPClient interface {
	GetMusicList(ctx context.Context, req *GetMusicListRequest, opts ...http.CallOption) (rsp *GetMusicListReply, err error)
	Ping(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type MusicHTTPClientImpl struct {
	cc *http.Client
}

func NewMusicHTTPClient(client *http.Client) MusicHTTPClient {
	return &MusicHTTPClientImpl{client}
}

func (c *MusicHTTPClientImpl) GetMusicList(ctx context.Context, in *GetMusicListRequest, opts ...http.CallOption) (*GetMusicListReply, error) {
	var out GetMusicListReply
	pattern := "/wxxcx/music/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/wxxcx.v1.music.Music/GetMusicList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MusicHTTPClientImpl) Ping(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/wxxcx/music/ping"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/wxxcx.v1.music.Music/Ping"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
