package biz

import (
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/medivhzhan/weapp/v3"
)

type WechatMpUseCase struct {
	log       *log.Helper
	clientMap map[string]*weapp.Client
	appMap    map[string]*conf.Application_App
}

func NewWechatMpUseCase(logger log.Logger, apps *conf.Application) *WechatMpUseCase {
	clientMap := make(map[string]*weapp.Client, len(apps.GetMp()))
	for _, app := range apps.GetMp() {
		if len(app.GetSecret()) == 0 {
			continue
		}
		clientMap[app.Id] = weapp.NewClient(app.GetId(), app.GetSecret())
	}
	return &WechatMpUseCase{log: log.NewHelper(logger), clientMap: clientMap, appMap: apps.GetMp()}
}

func (uc *WechatMpUseCase) GetApp(appid string) (*conf.Application_App, *weapp.Client) {
	return uc.appMap[appid], uc.clientMap[appid]
}
