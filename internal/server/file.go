package server

import (
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"net/http"
	"net/url"
	"strings"
)

func StripPrefix(logger log.Logger, cApp *conf.Application, prefix string, h http.Handler) http.Handler {
	if prefix == "" {
		return h
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			kind      string
			operation string
		)
		if info, ok := transport.FromServerContext(r.Context()); ok {
			kind = info.Kind().String()
			operation = info.Operation()
		}
		logContext := log.WithContext(r.Context(), logger)
		_ = logContext.Log(log.LevelInfo,
			"kind", "server",
			"component", kind,
			"operation", operation,
			"url", r.RequestURI,
		)
		if err := requestAuth(cApp, r, nil); err != nil {
			_ = logContext.Log(log.LevelError,
				"kind", "server",
				"component", kind,
				"operation", operation,
				"url", r.RequestURI,
				"referer", r.Referer(),
			)
			http.Error(w, `{"message":"非法请求",code:400}`, http.StatusOK)
			return
		}
		p := strings.TrimPrefix(r.URL.Path, prefix)
		rp := strings.TrimPrefix(r.URL.RawPath, prefix)

		if len(p) < len(r.URL.Path) && (r.URL.RawPath == "" || len(rp) < len(r.URL.RawPath)) {
			r2 := new(http.Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = p
			r2.URL.RawPath = rp
			h.ServeHTTP(w, r2)
		} else {
			http.NotFound(w, r)
		}
	})
}
