// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type WechatMpHTTPServer interface {
	LoginWechatMp(context.Context, *LoginWechatMpRequest) (*LoginWechatMpReply, error)
	SecurityCheckMsg(context.Context, *SecurityCheckMsgRequest) (*SecurityCheckMsgReply, error)
}

func RegisterWechatMpHTTPServer(s *http.Server, srv WechatMpHTTPServer) {
	r := s.Route("/")
	r.POST("/wxxcx/wechat/mp/login", _WechatMp_LoginWechatMp0_HTTP_Handler(srv))
	r.POST("/wxxcx/wechat/mp/sec-check/msg", _WechatMp_SecurityCheckMsg0_HTTP_Handler(srv))
}

func _WechatMp_LoginWechatMp0_HTTP_Handler(srv WechatMpHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginWechatMpRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.wxxcx.v1.WechatMp/LoginWechatMp")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginWechatMp(ctx, req.(*LoginWechatMpRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginWechatMpReply)
		return ctx.Result(200, reply)
	}
}

func _WechatMp_SecurityCheckMsg0_HTTP_Handler(srv WechatMpHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SecurityCheckMsgRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.wxxcx.v1.WechatMp/SecurityCheckMsg")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SecurityCheckMsg(ctx, req.(*SecurityCheckMsgRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SecurityCheckMsgReply)
		return ctx.Result(200, reply)
	}
}

type WechatMpHTTPClient interface {
	LoginWechatMp(ctx context.Context, req *LoginWechatMpRequest, opts ...http.CallOption) (rsp *LoginWechatMpReply, err error)
	SecurityCheckMsg(ctx context.Context, req *SecurityCheckMsgRequest, opts ...http.CallOption) (rsp *SecurityCheckMsgReply, err error)
}

type WechatMpHTTPClientImpl struct {
	cc *http.Client
}

func NewWechatMpHTTPClient(client *http.Client) WechatMpHTTPClient {
	return &WechatMpHTTPClientImpl{client}
}

func (c *WechatMpHTTPClientImpl) LoginWechatMp(ctx context.Context, in *LoginWechatMpRequest, opts ...http.CallOption) (*LoginWechatMpReply, error) {
	var out LoginWechatMpReply
	pattern := "/wxxcx/wechat/mp/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.wxxcx.v1.WechatMp/LoginWechatMp"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WechatMpHTTPClientImpl) SecurityCheckMsg(ctx context.Context, in *SecurityCheckMsgRequest, opts ...http.CallOption) (*SecurityCheckMsgReply, error) {
	var out SecurityCheckMsgReply
	pattern := "/wxxcx/wechat/mp/sec-check/msg"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.wxxcx.v1.WechatMp/SecurityCheckMsg"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
