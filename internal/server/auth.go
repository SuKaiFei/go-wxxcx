package server

import (
	"context"
	"encoding/json"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/SuKaiFei/go-wxxcx/internal/util"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	"strconv"
	"strings"
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
			if err = requestAuth(cApp, h.Request(), req); err != nil {
				return nil, err
			}

			return handler(ctx, req)
		}
	}
}

func requestAuth(cApp *conf.Application, request *http.Request, req interface{}) (err error) {
	referer := request.Header.Get("Referer")
	requestUrl := request.URL
	if len(referer) < 44 {
		return ErrBadAppid
	}
	appid := referer[26:44]
	app, found := cApp.GetMp()[appid]
	if !found {
		return ErrBadAppid
	}

	var (
		data         map[string]interface{}
		reqSign      string
		timestampStr string
	)
	if strings.ToLower(request.Method) == "get" {
		timestamps := requestUrl.Query()["timestamp"]
		if len(timestamps) == 0 || len(timestamps[0]) == 0 {
			return ErrBadSign
		}
		timestampStr = timestamps[0]
		if err != nil {
			return ErrBadSign
		}

		data = make(map[string]interface{}, len(requestUrl.Query()))
		for k, v := range requestUrl.Query() {
			data[k] = v[0]
		}
		reqSign = requestUrl.Query()["sign"][0]
	} else {
		bytes, err := json.Marshal(req)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(bytes, &data); err != nil {
			return err
		}

		timestampStr = data["timestamp"].(string)
		if signObj, found := data["sign"]; found {
			reqSign = signObj.(string)
		}
	}

	timestamp, _ := strconv.Atoi(timestampStr)
	if time.Since(time.Unix(int64(timestamp), 0)) > 10*time.Second {
		return ErrBadSign
	}

	sign := util.GetSign(data, app.GetKey())
	if reqSign != sign {
		return ErrBadSign
	}

	return nil
}
