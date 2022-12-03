package biz

import (
	"context"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

type WechatOcUseCase struct {
	log       *log.Helper
	clientMap map[string]*officialaccount.OfficialAccount
	appMap    map[string]*conf.Application_App
}

func NewWechatOcUseCase(logger log.Logger, apps *conf.Application, dataConf *conf.Data) *WechatOcUseCase {
	clientMap := make(map[string]*officialaccount.OfficialAccount, len(apps.GetMp()))
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
		if len(app.GetSecret()) == 0 || len(app.GetKey()) > 0 {
			continue
		}
		clientMap[app.Id] = wc.GetOfficialAccount(&offConfig.Config{
			AppID:          app.GetId(),
			AppSecret:      app.GetSecret(),
			Token:          app.GetToken(),
			EncodingAESKey: app.GetEncodingAESKey(),
		})
	}
	return &WechatOcUseCase{log: log.NewHelper(logger), clientMap: clientMap, appMap: apps.GetMp()}
}

func (uc *WechatOcUseCase) GetApp(appid string) (*conf.Application_App, *officialaccount.OfficialAccount) {
	return uc.appMap[appid], uc.clientMap[appid]
}
