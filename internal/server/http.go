package server

import (
	"bytes"
	"context"
	"fmt"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	"github.com/go-kratos/kratos/v2/errors"
	"io"
	nhttp "net/http"

	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/SuKaiFei/go-wxxcx/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var whiteList = map[string]struct{}{
	"/wxxcx.v1.bqb.Bqb/Ping": {},
	//"/api.wxxcx.v1.Image/UploadImage": {},
}

func NewWhiteListMatcher() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, cApp *conf.Application,
	bqbSVC *service.BqbService,
	voiceSVC *service.VoiceService,
	articleSVC *service.ArticleService,
	navigationSVC *service.NavigationService,
	wechatMpSVC *service.WechatMpService,
	wechatOcSVC *service.WechatOcService,
	imageSVC *service.ImageService,
	musicSVC *service.MusicService,
	chatGptSVC *service.ChatGptService,
	communitySVC *service.CommunityService,
	wordcloudSVC *service.WordcloudService,
	securityUC *biz.SecurityUseCase,
	logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			selector.Server(MiddlewareAuth(cApp, securityUC)).
				Match(NewWhiteListMatcher()).
				Build(),
		),
		http.RequestDecoder(func(r *nhttp.Request, v interface{}) error {
			codec, ok := http.CodecForRequest(r, "Content-Type")
			if !ok {
				return errors.BadRequest("CODEC", fmt.Sprintf("unregister Content-Type: %s", r.Header.Get("Content-Type")))
			}
			data, err := io.ReadAll(r.Body)
			if r.URL.Path != "/wxxcx/image/upload" {
				r.Body = io.NopCloser(bytes.NewBuffer(data))
			}
			if err != nil {
				return errors.BadRequest("CODEC", err.Error())
			}
			if len(data) == 0 {
				return nil
			}
			if err = codec.Unmarshal(data, v); err != nil {
				return errors.BadRequest("CODEC", fmt.Sprintf("body unmarshal %s", err.Error()))
			}
			return nil
		}),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	fs := nhttp.FileServer(nhttp.Dir(cApp.GetStaticPath()))
	srv.HandlePrefix("/static/", StripPrefix(logger, cApp, "/static/", fs, securityUC))

	v1.RegisterBqbHTTPServer(srv, bqbSVC)
	v1.RegisterVoiceHTTPServer(srv, voiceSVC)
	v1.RegisterArticleHTTPServer(srv, articleSVC)
	v1.RegisterNavigationHTTPServer(srv, navigationSVC)
	v1.RegisterWechatMpHTTPServer(srv, wechatMpSVC)
	v1.RegisterImageHTTPServer(srv, imageSVC)
	v1.RegisterMusicHTTPServer(srv, musicSVC)
	v1.RegisterChatGptHTTPServer(srv, chatGptSVC)
	v1.RegisterCommunityHTTPServer(srv, communitySVC)
	v1.RegisterWordcloudHTTPServer(srv, wordcloudSVC)
	RegisterWechatHTTPServer(srv, wechatMpSVC)
	return srv
}
