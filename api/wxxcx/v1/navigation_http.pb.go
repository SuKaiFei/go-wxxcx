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

type NavigationHTTPServer interface {
	GetNavigations(context.Context, *GetNavigationsRequest) (*GetNavigationsReply, error)
}

func RegisterNavigationHTTPServer(s *http.Server, srv NavigationHTTPServer) {
	r := s.Route("/")
	r.GET("/wxxcx/navigation/list_by_code", _Navigation_GetNavigations0_HTTP_Handler(srv))
}

func _Navigation_GetNavigations0_HTTP_Handler(srv NavigationHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetNavigationsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/wxxcx.v1.navigation.Navigation/GetNavigations")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNavigations(ctx, req.(*GetNavigationsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetNavigationsReply)
		return ctx.Result(200, reply)
	}
}

type NavigationHTTPClient interface {
	GetNavigations(ctx context.Context, req *GetNavigationsRequest, opts ...http.CallOption) (rsp *GetNavigationsReply, err error)
}

type NavigationHTTPClientImpl struct {
	cc *http.Client
}

func NewNavigationHTTPClient(client *http.Client) NavigationHTTPClient {
	return &NavigationHTTPClientImpl{client}
}

func (c *NavigationHTTPClientImpl) GetNavigations(ctx context.Context, in *GetNavigationsRequest, opts ...http.CallOption) (*GetNavigationsReply, error) {
	var out GetNavigationsReply
	pattern := "/wxxcx/navigation/list_by_code"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/wxxcx.v1.navigation.Navigation/GetNavigations"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
