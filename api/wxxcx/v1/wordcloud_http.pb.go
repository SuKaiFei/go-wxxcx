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

type WordcloudHTTPServer interface {
	GenerateWordcloudImage(context.Context, *GenerateWordcloudImageRequest) (*UploadImageReply, error)
	Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
}

func RegisterWordcloudHTTPServer(s *http.Server, srv WordcloudHTTPServer) {
	r := s.Route("/")
	r.POST("/wxxcx/wordcloud/generate", _Wordcloud_GenerateWordcloudImage0_HTTP_Handler(srv))
	r.GET("/wxxcx/wordcloud/ping", _Wordcloud_Ping3_HTTP_Handler(srv))
}

func _Wordcloud_GenerateWordcloudImage0_HTTP_Handler(srv WordcloudHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GenerateWordcloudImageRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/wxxcx.v1.wordcloud.Wordcloud/GenerateWordcloudImage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GenerateWordcloudImage(ctx, req.(*GenerateWordcloudImageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UploadImageReply)
		return ctx.Result(200, reply)
	}
}

func _Wordcloud_Ping3_HTTP_Handler(srv WordcloudHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/wxxcx.v1.wordcloud.Wordcloud/Ping")
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

type WordcloudHTTPClient interface {
	GenerateWordcloudImage(ctx context.Context, req *GenerateWordcloudImageRequest, opts ...http.CallOption) (rsp *UploadImageReply, err error)
	Ping(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type WordcloudHTTPClientImpl struct {
	cc *http.Client
}

func NewWordcloudHTTPClient(client *http.Client) WordcloudHTTPClient {
	return &WordcloudHTTPClientImpl{client}
}

func (c *WordcloudHTTPClientImpl) GenerateWordcloudImage(ctx context.Context, in *GenerateWordcloudImageRequest, opts ...http.CallOption) (*UploadImageReply, error) {
	var out UploadImageReply
	pattern := "/wxxcx/wordcloud/generate"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/wxxcx.v1.wordcloud.Wordcloud/GenerateWordcloudImage"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WordcloudHTTPClientImpl) Ping(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/wxxcx/wordcloud/ping"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/wxxcx.v1.wordcloud.Wordcloud/Ping"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
