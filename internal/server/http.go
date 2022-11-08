package server

import (
	"context"
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
	"/wxxcx.v1.Bqb/Ping": {},
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
	logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			selector.Server(MiddlewareAuth(cApp)).
				Match(NewWhiteListMatcher()).
				Build(),
		),
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
	srv.HandlePrefix("/static/", StripPrefix(logger, cApp, "/static/", fs))

	v1.RegisterBqbHTTPServer(srv, bqbSVC)
	v1.RegisterVoiceHTTPServer(srv, voiceSVC)
	v1.RegisterArticleHTTPServer(srv, articleSVC)
	v1.RegisterNavigationHTTPServer(srv, navigationSVC)
	return srv
}
