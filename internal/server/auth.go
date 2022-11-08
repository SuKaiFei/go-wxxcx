package server

import (
	"context"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/SuKaiFei/go-wxxcx/internal/util"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	"net/url"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
)

var (
	ErrBadAppid = errors.BadRequest("请求异常", "appid错误")
	ErrBadSign  = errors.BadRequest("请求异常", "sign错误")
)

func MiddlewareAuth(cApp *conf.Application) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			h := ctx.(http.Context)
			request := h.Request()
			if err = requestAuth(cApp, request.Header.Get("Referer"), request.URL); err != nil {
				return nil, err
			}

			return handler(ctx, req)
		}
	}
}

func requestAuth(cApp *conf.Application, referer string, requestUrl *url.URL) error {
	if len(referer) < 44 {
		return ErrBadAppid
	}
	appid := referer[26:44]
	app, found := cApp.GetMp()[appid]
	if !found {
		return ErrBadAppid
	}

	timestamps := requestUrl.Query()["timestamp"]
	if len(timestamps) == 0 || len(timestamps[0]) == 0 {
		return ErrBadSign
	}
	timestamp, err := strconv.Atoi(timestamps[0])
	if err != nil {
		return ErrBadSign
	}
	if time.Since(time.Unix(int64(timestamp), 0)) > 5*time.Second {
		return ErrBadSign
	}

	sign := util.GetSign(requestUrl.Query(), app.GetKey())
	if requestUrl.Query()["sign"][0] != sign {
		return ErrBadSign
	}

	return nil
}
