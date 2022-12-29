//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package service

import (
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/SuKaiFei/go-wxxcx/internal/data"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

type unitTestSvc struct {
	bqbSvc        *BqbService
	articleSvc    *ArticleService
	voiceSvc      *VoiceService
	navigationSvc *NavigationService
	wechatMpSvc   *WechatMpService
	wechatOcSvc   *WechatOcService
	imageSvc      *ImageService
	musicSvc      *MusicService
	chatGptSvc    *ChatGptService
	communitySvc  *CommunityService
}

func NewTestUnitTestSvcService(*conf.Server, log.Logger, *conf.Bootstrap, *conf.Data, *conf.Application) (*unitTestSvc, func(), error) {
	panic(wire.Build(
		wire.Struct(new(unitTestSvc), "*"),
		data.ProviderSet,
		biz.ProviderSet,
		ProviderSet),
	)
}
