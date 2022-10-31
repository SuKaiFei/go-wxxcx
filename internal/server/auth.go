package server

import (
	"context"
	"github.com/SuKaiFei/go-wxxcx/internal/util"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	"strconv"
	"time"

	"github.com/SuKaiFei/go-wxxcx/internal/conf"

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

			referer := request.Header.Get("Referer")
			if len(referer) < 44 {
				return nil, ErrBadAppid
			}
			appid := referer[26:44]
			app, found := cApp.GetMp()[appid]
			if !found {
				return nil, ErrBadAppid
			}

			timestamps := request.URL.Query()["timestamp"]
			if len(timestamps) == 0 || len(timestamps[0]) == 0 {
				return nil, ErrBadSign
			}
			timestamp, err := strconv.Atoi(timestamps[0])
			if err != nil {
				return nil, ErrBadSign
			}
			if time.Since(time.Unix(int64(timestamp), 0)) > 5*time.Second {
				return nil, ErrBadSign
			}

			sign := util.GetSign(request.URL.Query(), app.GetKey())
			if request.URL.Query()["sign"][0] != sign {
				return nil, ErrBadSign
			}

			return handler(ctx, req)
		}
	}
}
