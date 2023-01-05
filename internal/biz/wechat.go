package biz

import (
	"context"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	errors2 "github.com/pkg/errors"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/miniprogram/security"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

type WechatRepo interface {
	GetRedisClient() redis.UniversalClient
	UpsertUser(context.Context, *WechatUser) error
	GetUser(context.Context, string, string) (*WechatUser, error)
}

type WechatUseCase struct {
	log         *log.Helper
	repo        WechatRepo
	oaClientMap map[string]*officialaccount.OfficialAccount
	mpClientMap map[string]*miniprogram.MiniProgram
	appMap      map[string]*conf.Application_App
}

func NewWechatUseCase(logger log.Logger, apps *conf.Application, repo WechatRepo) *WechatUseCase {
	oaClientMap := make(map[string]*officialaccount.OfficialAccount, len(apps.GetMp()))
	mpClientMap := make(map[string]*miniprogram.MiniProgram, len(apps.GetMp()))
	wc := wechat.NewWechat()
	redis := new(cache.Redis)
	redis.SetConn(repo.GetRedisClient())
	redis.SetRedisCtx(context.Background())
	wc.SetCache(redis)
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
		repo:        repo,
	}
}

func (uc *WechatUseCase) MsgCheck(appid, openid, content string) (*security.MsgCheckResponse, error) {
	if len(content) == 0 {
		res := new(security.MsgCheckResponse)
		res.Result.Label = 100
		return res, nil
	}

	client := uc.GetMpApp(appid)
	securityCli := client.GetSecurity()
	securityRequest := &security.MsgCheckRequest{
		Scene:   2,
		OpenID:  openid,
		Content: content,
	}
	secCheck, err := securityCli.MsgCheck(securityRequest)
	if err != nil {
		return nil, err
	}
	if secCheck.ErrCode != 0 {
		return nil, errors2.New(secCheck.Error())
	}
	if secCheck.Result.Label != 100 {
		log.Warnw(
			"message", "SecurityCheckMsg Label!=100",
			"content", content,
			"appid", appid,
			"openid", openid,
		)
	}
	return &secCheck, nil
}

func (uc *WechatUseCase) Code2Session(appid, code string) (*auth.ResCode2Session, error) {
	session, err := uc.mpClientMap[appid].GetAuth().Code2Session(code)
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	if session.ErrCode != 0 {
		return nil, errors2.New(session.Error())
	}
	go func() {
		user := &WechatUser{
			Appid:   appid,
			Openid:  session.OpenID,
			Unionid: session.UnionID,
		}
		if err := uc.repo.UpsertUser(context.Background(), user); err != nil {
			uc.log.Errorw("msg", "repo.UpsertUser", "err", err)
			return
		}
	}()
	return &session, nil
}

func (uc *WechatUseCase) GetOaApp(appid string) *officialaccount.OfficialAccount {
	return uc.oaClientMap[appid]
}

func (uc *WechatUseCase) GetMpApp(appid string) *miniprogram.MiniProgram {
	return uc.mpClientMap[appid]
}

func (uc *WechatUseCase) GetUser(ctx context.Context, appid, openid string) (*WechatUser, error) {
	user, err := uc.repo.GetUser(ctx, appid, openid)
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return user, nil
}

func (uc *WechatUseCase) SyncOAUser(ctx context.Context, appid, nextOpenid string) error {
	app := uc.GetOaApp(appid)
	if app == nil {
		return errors2.New("app is empty")
	}

	openidList, err := app.GetUser().ListUserOpenIDs(nextOpenid)
	if err != nil {
		return errors2.WithStack(err)
	}
	if openidList.ErrCode != 0 {
		return errors2.New(openidList.Error())
	}
	for _, openID := range openidList.Data.OpenIDs {
		userInfo, err := app.GetUser().GetUserInfo(openID)
		if err != nil {
			return errors2.WithStack(err)
		}
		_ = userInfo
	}

	return nil
}
