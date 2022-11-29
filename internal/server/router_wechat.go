package server

import (
	"fmt"
	"github.com/SuKaiFei/go-wxxcx/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"io"
)

func RegisterWechatHTTPServer(s *http.Server, wechatMpSVC *service.WechatMpService) {
	r := s.Route("/")
	r.GET("/wechat/miniapp/notify/{appid}", wechatMiniappNotifyHandle(wechatMpSVC))
}

func wechatMiniappNotifyHandle(wechatMpSVC *service.WechatMpService) func(ctx http.Context) error {
	//  通用处理器
	handler := func(req map[string]interface{}) map[string]interface{} {
		log.Infow(
			"msg", "wechatMiniappNotifyHandle通用处理器",
			"req", fmt.Sprintf("%+v", req),
		)
		switch req["MsgType"] {
		case "text":
			// Do something cool ...
		}

		return nil
	}
	return func(ctx http.Context) error {
		appid := ctx.Request().RequestURI[23 : 23+18]
		confApp, sdk := wechatMpSVC.GetApp(appid)
		body := ctx.Request().Body
		bodyBytes, err := io.ReadAll(body)
		if err != nil {
			return ctx.Result(400, nil)
		}
		log.Infow(
			"msg", "wechatMiniappNotifyHandle",
			"BodyBytes", string(bodyBytes),
			"ctx.Query().Encode()", ctx.Query().Encode(),
			"ctx.Form().Encode()", ctx.Form().Encode(),
			"appid", appid,
			"sdkIsNil", sdk == nil,
		)
		srv, err := sdk.NewServer(
			confApp.GetToken(),
			confApp.GetEncodingAESKey(),
			"mchID",
			"apiKey",
			true,
			handler,
		)
		if err != nil {
			log.Errorw("init server error: %s", err)
			return err
		}

		if err := srv.Serve(ctx.Response(), ctx.Request()); err != nil {
			log.Errorw("serving error: %s", err)
			return err
		}
		return nil
	}
}
