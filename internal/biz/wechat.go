package biz

import (
	"context"
	"fmt"
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
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/silenceper/wechat/v2/officialaccount/user"
	"gorm.io/gorm"
	"strings"
	"sync"
	"time"
)

type WechatRepo interface {
	GetRedisClient() redis.UniversalClient
	UpsertUser(context.Context, *WechatUser) error
	GetUser(context.Context, string, string) (*WechatUser, error)
	GetOAUser(context.Context, string) (*WechatOAUser, error)
	UpsertOAUser(context.Context, []*WechatOAUser) error
	GetOAUserByOpenid(context.Context, []string) ([]*WechatOAUser, error)
}

type WechatUseCase struct {
	log         *log.Helper
	repo        WechatRepo
	oaClientMap map[string]*officialaccount.OfficialAccount
	mpClientMap map[string]*miniprogram.MiniProgram
	appMap      map[string]*conf.Application_App

	//sendCH   chan string
	//sendTask *sync.WaitGroup
}

func NewWechatUseCase(logger log.Logger, apps *conf.Application, repo WechatRepo) (*WechatUseCase, func(), error) {
	const sendCHCount = 50
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

	uc := &WechatUseCase{
		log:         log.NewHelper(logger),
		oaClientMap: oaClientMap,
		mpClientMap: mpClientMap,
		appMap:      apps.GetMp(),
		repo:        repo,
	}
	//sendCH := make(chan string, sendCHCount)
	//sendTask := new(sync.WaitGroup)
	//for i := 0; i < sendCHCount; i++ {
	//	uc.sendTask.Add(1)
	//	go func() {
	//		defer uc.sendTask.Done()
	//		for openid := range uc.sendCH {
	//			msgID, err := uc.Send(openid)
	//			log.Infof("index(%d) msgID(%d) error(%+v)\n", i, msgID, err)
	//		}
	//	}()
	//}
	ctx, cancelFunc := context.WithCancel(context.Background())
	go uc.syncOAUserTask(ctx, "wx9ef62ba2e3525812")

	closeFunc := func() {
		//close(uc.sendCH)
		//uc.sendTask.Wait()
		cancelFunc()
	}
	return uc, closeFunc, nil
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
			"appidCommunity", appid,
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

func (uc *WechatUseCase) SendTemplateMsg(ctx context.Context, unionid string, msg *message.TemplateMessage) (int64, error) {
	oaUser, err := uc.repo.GetOAUser(ctx, unionid)
	if err != nil {
		if errors2.Cause(err) == gorm.ErrRecordNotFound {
			return 0, errors2.New("没有关注公众号")
		}
		return 0, errors2.WithStack(err)
	}

	app := uc.GetOaApp(appidOA)
	if app == nil {
		return 0, errors2.New("app is empty")
	}

	msg.ToUser = oaUser.Openid
	msgID, err := app.GetTemplate().Send(msg)
	if err != nil {
		return 0, errors2.WithStack(err)
	}

	return msgID, nil
}

func (uc *WechatUseCase) GetUser(ctx context.Context, appid, openid string) (*WechatUser, error) {
	user, err := uc.repo.GetUser(ctx, appid, openid)
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return user, nil
}

func (uc *WechatUseCase) syncOAUserTask(ctx context.Context, appid string) {
	index := 0
	nextOpenid := ""
	for {
		time.Sleep(time.Minute * 45)
		for {
			index++
			count, nextOpenidTmp, err := uc.syncOAUser(ctx, appid, nextOpenid)
			if err != nil {
				uc.log.Errorw("msg", "syncOAUserTask", "err", err)
				break
			}
			nextOpenid = nextOpenidTmp
			uc.log.Infow("msg", "syncOAUserTask", "count", count, "index", index, "nextOpenid", nextOpenidTmp)
			if count < 10000 {
				uc.log.Infow("msg", "SyncOAUserTaskDone", "index", index)
				nextOpenid = ""
				index = 0
				break
			}
		}
		//time.Sleep(time.Minute * 45)
	}
}

func (uc *WechatUseCase) syncOAUser(ctx context.Context, appid, nextOpenid string) (int, string, error) {
	app := uc.GetOaApp(appid)
	if app == nil {
		return 0, "", errors2.New("app is empty")
	}

	openidList, err := app.GetUser().ListUserOpenIDs(nextOpenid)
	if err != nil {
		return 0, "", errors2.WithStack(err)
	}
	if openidList.ErrCode != 0 {
		return 0, "", errors2.New(openidList.Error())
	}
	openIDs := openidList.Data.OpenIDs
	emptyOpenidList := make([]string, 0, len(openIDs))
	oaUserMap := make(map[string]struct{}, len(openIDs))

	oaUserList, err := uc.repo.GetOAUserByOpenid(ctx, openIDs)
	for _, oaUser := range oaUserList {
		oaUserMap[oaUser.Openid] = struct{}{}
	}

	for _, openid := range openIDs {
		if _, found := oaUserMap[openid]; found {
			continue
		}
		emptyOpenidList = append(emptyOpenidList, openid)
	}

	var (
		errUserinfo           error
		cancelCtx, cancelFunc = context.WithCancel(ctx)
		userinfoList          = make([]*user.Info, 0, len(emptyOpenidList))
		wg                    = new(sync.WaitGroup)
		lock                  = new(sync.Mutex)
	)
	defer cancelFunc()

	for i, openID := range emptyOpenidList {
		wg.Add(1)
		go func(i int, openID string) {
			defer wg.Done()

			select {
			case <-cancelCtx.Done():
			default:
				userinfo, err := app.GetUser().GetUserInfo(openID)
				lock.Lock()
				if err != nil {
					errUserinfo = err
					uc.log.Errorw(
						"msg", "SyncOAUserGetUserinfo",
						"err", err,
					)
					cancelFunc()
				}
				userinfoList = append(userinfoList, userinfo)
				lock.Unlock()
			}
		}(i, openID)
	}
	wg.Wait()
	if errUserinfo != nil {
		return 0, "", errors2.WithStack(errUserinfo)
	}

	oauserList := make([]*WechatOAUser, len(emptyOpenidList))
	for i, info := range userinfoList {
		oauserList[i] = &WechatOAUser{
			Openid:         info.OpenID,
			Unionid:        info.UnionID,
			Subscribe:      info.Subscribe,
			Nickname:       info.Nickname,
			Sex:            info.Sex,
			City:           info.City,
			Country:        info.Country,
			Province:       info.Province,
			Language:       info.Language,
			Headimgurl:     info.Headimgurl,
			SubscribeTime:  info.SubscribeTime,
			Remark:         info.Remark,
			GroupID:        info.GroupID,
			SubscribeScene: info.SubscribeScene,
			QrScene:        info.QrScene,
			QrSceneStr:     info.QrSceneStr,
		}
		if len(info.TagIDList) > 0 {
			oauserList[i].TagIDList = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(info.TagIDList)), ","), "[]")
		}
	}
	if err := uc.repo.UpsertOAUser(ctx, oauserList); err != nil {
		uc.log.Errorw("msg", "repo.UpsertOAUser", "err", err)
		return 0, "", errors2.WithStack(err)
	}

	return len(openIDs), openIDs[len(openIDs)-1], nil
}
