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

type ImageHTTPServer interface {
	Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	UploadImage(context.Context, *UploadImageRequest) (*UploadImageReply, error)
}

func RegisterImageHTTPServer(s *http.Server, srv ImageHTTPServer) {
	r := s.Route("/")
	r.POST("/wxxcx/image/upload", _Image_UploadImage0_HTTP_Handler(srv))
	r.GET("/wxxcx/image/ping", _Image_Ping3_HTTP_Handler(srv))
}

func _Image_UploadImage0_HTTP_Handler(srv ImageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UploadImageRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.wxxcx.v1.Image/UploadImage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadImage(ctx, req.(*UploadImageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UploadImageReply)
		return ctx.Result(200, reply)
	}
}

func _Image_Ping3_HTTP_Handler(srv ImageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.wxxcx.v1.Image/Ping")
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

type ImageHTTPClient interface {
	Ping(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	UploadImage(ctx context.Context, req *UploadImageRequest, opts ...http.CallOption) (rsp *UploadImageReply, err error)
}

type ImageHTTPClientImpl struct {
	cc *http.Client
}

func NewImageHTTPClient(client *http.Client) ImageHTTPClient {
	return &ImageHTTPClientImpl{client}
}

func (c *ImageHTTPClientImpl) Ping(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/wxxcx/image/ping"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.wxxcx.v1.Image/Ping"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ImageHTTPClientImpl) UploadImage(ctx context.Context, in *UploadImageRequest, opts ...http.CallOption) (*UploadImageReply, error) {
	var out UploadImageReply
	pattern := "/wxxcx/image/upload"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.wxxcx.v1.Image/UploadImage"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
