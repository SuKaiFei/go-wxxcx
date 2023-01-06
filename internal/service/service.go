package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewBqbService,
	NewVoiceService,
	NewArticleService,
	NewNavigationService,
	NewWechatMpService,
	NewImageService,
	NewMusicService,
	NewWechatOAService,
	NewChatGptService,
	NewCommunityService,
	NewWordcloudService,
)
