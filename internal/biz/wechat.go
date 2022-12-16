package biz

import (
	"context"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

type WechatUseCase struct {
	log         *log.Helper
	oaClientMap map[string]*officialaccount.OfficialAccount
	mpClientMap map[string]*miniprogram.MiniProgram
	appMap      map[string]*conf.Application_App
}

func NewWechatUseCase(logger log.Logger, apps *conf.Application, dataConf *conf.Data) *WechatUseCase {
	oaClientMap := make(map[string]*officialaccount.OfficialAccount, len(apps.GetMp()))
	mpClientMap := make(map[string]*miniprogram.MiniProgram, len(apps.GetMp()))
	wc := wechat.NewWechat()
	wc.SetCache(cache.NewRedis(context.Background(), &cache.RedisOpts{
		Host:        dataConf.Redis.GetAddr(),
		Password:    dataConf.Redis.GetPassword(),
		Database:    0,
		MaxIdle:     10000,
		MaxActive:   10000,
		IdleTimeout: 60,
	}))
	for _, app := range apps.GetMp() {
		if len(app.GetSecret()) == 0 {
			continue
		}
		if len(app.GetKey()) > 0 {
			mpClientMap[app.Id] = wc.GetMiniProgram(&miniConfig.Config{
				AppID:     app.GetId(),
				AppSecret: app.GetSecret(),
			})
		} else {
			oaClientMap[app.Id] = wc.GetOfficialAccount(&offConfig.Config{
				AppID:          app.GetId(),
				AppSecret:      app.GetSecret(),
				Token:          app.GetToken(),
				EncodingAESKey: app.GetEncodingAESKey(),
			})
		}
	}
	return &WechatUseCase{
		log:         log.NewHelper(logger),
		oaClientMap: oaClientMap,
		mpClientMap: mpClientMap,
		appMap:      apps.GetMp(),
	}
}

func (uc *WechatUseCase) GetOaApp(appid string) *officialaccount.OfficialAccount {
	return uc.oaClientMap[appid]
}

func (uc *WechatUseCase) GetMpApp(appid string) *miniprogram.MiniProgram {
	return uc.mpClientMap[appid]
}
