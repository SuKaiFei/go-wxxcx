// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/SuKaiFei/go-wxxcx/internal/data"
	"github.com/SuKaiFei/go-wxxcx/internal/server"
	"github.com/SuKaiFei/go-wxxcx/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, application *conf.Application, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	biaoQingBaoRepo := data.NewBqbRepo(dataData, logger)
	biaoQingBaoUseCase := biz.NewBiaoQingBaoUseCase(biaoQingBaoRepo, logger)
	bqbService := service.NewBqbService(biaoQingBaoUseCase)
	voiceRepo := data.NewVoiceRepo(dataData, logger)
	voiceUseCase := biz.NewVoiceUseCase(voiceRepo, logger)
	voiceService := service.NewVoiceService(voiceUseCase)
	articleRepo := data.NewArticleRepo(dataData, logger)
	articleUseCase := biz.NewArticleUseCase(articleRepo, logger)
	articleService := service.NewArticleService(articleUseCase)
	navigationRepo := data.NewNavigationRepo(dataData, logger)
	navigationUseCase := biz.NewNavigationUseCase(navigationRepo, logger)
	navigationService := service.NewNavigationService(navigationUseCase)
	wechatRepo := data.NewWechatRepo(dataData, logger)
	wechatUseCase := biz.NewWechatUseCase(logger, application, wechatRepo)
	communityRepo := data.NewCommunityRepo(dataData, logger)
	cosUseCase := biz.NewCosUseCase(application, logger)
	communityUseCase := biz.NewCommunityUseCase(communityRepo, cosUseCase, logger)
	wechatMpService := service.NewWechatMpService(wechatUseCase, communityUseCase)
	wechatOcService, cleanup2, err := service.NewWechatOcService(wechatUseCase)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	imageUseCase := biz.NewImageUseCase(logger)
	imageService := service.NewImageService(imageUseCase, cosUseCase, application)
	musicRepo := data.NewMusicRepo(dataData, logger)
	musicUseCase := biz.NewMusicUseCase(musicRepo, logger)
	musicService := service.NewMusicService(musicUseCase)
	chatGPTRepo := data.NewChatGPTRepo(dataData, logger)
	chatGPTUseCase := biz.NewChatGPTUseCase(chatGPTRepo, application, logger)
	chatGptService := service.NewChatGptService(chatGPTUseCase)
	communityService := service.NewCommunityService(communityUseCase, cosUseCase, wechatUseCase)
	securityRepo := data.NewSecurityRepo(dataData, logger)
	securityUseCase := biz.NewSecurityUseCase(securityRepo, logger)
	httpServer := server.NewHTTPServer(confServer, application, bqbService, voiceService, articleService, navigationService, wechatMpService, wechatOcService, imageService, musicService, chatGptService, communityService, securityUseCase, logger)
	app := newApp(logger, httpServer)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
