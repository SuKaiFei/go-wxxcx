package server

import (
	"context"
	"fmt"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/SuKaiFei/go-wxxcx/internal/util"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
)

var (
	ErrBadAppid = errors.BadRequest("请求异常", "appid错误")
	ErrBadSign  = errors.BadRequest("请求异常", "sign错误")
)

func MiddlewareAuth(cApp *conf.Application, securityUC *biz.SecurityUseCase) middleware.Middleware {

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			h := ctx.(http.Context)
			if err = requestAuth(cApp, h.Request(), req, securityUC); err != nil {
				return nil, err
			}

			return handler(ctx, req)
		}
	}
}

func requestAuth(cApp *conf.Application, request *http.Request, req interface{}, securityUC *biz.SecurityUseCase) (err error) {
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
		openid       string
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
		if len(requestUrl.Query()["openid"]) > 0 && len(requestUrl.Query()["openid"][0]) > 0 {
			openid = requestUrl.Query()["openid"][0]
		}
	} else {
		bytes, err := jsoniter.Marshal(req)
		if err != nil {
			return err
		}
		if err := jsoniter.Unmarshal(bytes, &data); err != nil {
			return err
		}

		timestampStr = data["timestamp"].(string)
		if signObj, found := data["sign"]; found {
			reqSign = signObj.(string)
		}
		if openidTmp, found := data["openid"]; found {
			openid = openidTmp.(string)
		}

		delete(data, "file")
	}

	timestamp, _ := strconv.Atoi(timestampStr)
	var (
		reqTime time.Time
	)

	if len(timestampStr) == 13 && openid != "" {
		reqTime = time.Unix(int64(timestamp/1000), 0)
	} else {
		reqTime = time.Unix(int64(timestamp), 0)
	}
	if time.Since(reqTime) > 5*time.Minute {
		log.Errorw(
			"msg", fmt.Sprintf("时间戳超时%d", time.Now().Unix()-int64(timestamp)),
			"请求IP", request.Header.Get("X-Forward-For"),
			"请求地址", request.RequestURI,
			"UserAgent", request.UserAgent(),
		)
		return ErrBadSign
	}

	sign := util.GetSign(data, app.GetKey())
	if reqSign != sign {
		return ErrBadSign
	}

	if len(timestampStr) == 13 && len(openid) > 0 {
		if err = securityUC.VerifySign(request.Context(), requestUrl.RequestURI(), appid, openid, sign); err != nil {
			log.Errorw(
				"msg", "VerifySign",
				"请求IP", request.Header.Get("X-Forward-For"),
				"请求地址", request.RequestURI,
				"UserAgent", request.UserAgent(),
				"sign", sign,
				"err", fmt.Sprintf("%+v", err),
			)
			return ErrBadSign
		}
	}

	return nil
}
