package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewBiaoQingBaoUseCase,
	NewVoiceUseCase,
	NewArticleUseCase,
	NewNavigationUseCase,
	NewImageUseCase,
	NewMusicUseCase,
	NewWechatUseCase,
	NewChatGPTUseCase,
	NewCommunityUseCase,
	NewCosUseCase,
	NewSecurityUseCase,
	NewWordcloudUseCase,
)

const (
	AppidCommunity = "wxec615f70feb4e93c"
	appidOA        = "wx9ef62ba2e3525812"
)
