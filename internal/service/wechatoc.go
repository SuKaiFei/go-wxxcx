package service

import (
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/silenceper/wechat/v2/officialaccount/user"
	"sync"
)

type WechatOcService struct {
	uc       *biz.WechatOcUseCase
	sendCH   chan string
	sendTask *sync.WaitGroup
}

func NewWechatOcService(uc *biz.WechatOcUseCase) (*WechatOcService, func(), error) {
	const sendCHCount = 50
	s := &WechatOcService{
		uc:       uc,
		sendCH:   make(chan string, sendCHCount),
		sendTask: new(sync.WaitGroup),
	}

	for i := 0; i < sendCHCount; i++ {
		s.sendTask.Add(1)
		go func() {
			defer s.sendTask.Done()
			for openid := range s.sendCH {
				msgID, err := s.Send(openid)
				log.Infof("index(%d) msgID(%d) error(%+v)\n", i, msgID, err)
			}
		}()
	}
	closeFunc := func() {
		close(s.sendCH)
		s.sendTask.Wait()
		log.Info("NewWechatOcService closeFunc")
	}
	return s, closeFunc, nil
}

func (s *WechatOcService) GetUserList(appid string) (
	[]*user.Info,
	error,
) {
	_, client := s.GetApp(appid)
	userCli := client.GetUser()
	tags, err := userCli.GetTag()
	if err != nil {
		return nil, err
	}

	var (
		userOpenidList []string
	)

	for _, tag := range tags {
		userListTmp, err := userCli.OpenIDListByTag(tag.ID, "")
		if err != nil {
			return nil, err
		}
		userOpenidList = append(userOpenidList, userListTmp.Data.OpenIDs...)
	}

	userinfoList := make([]*user.Info, len(userOpenidList))
	for i, openid := range userOpenidList {
		info, err := userCli.GetUserInfo(openid)
		if err != nil {
			return nil, err
		}
		userinfoList[i] = info
	}

	return userinfoList, nil
}

func (s *WechatOcService) SendAsync(openid string) {
	s.sendCH <- openid
}

func (s *WechatOcService) Send(openid string) (int64, error) {
	const templateID = "3EHnd170OYZ7ucWq1cOdRdeGGzrXD0u73QkTYHZHc4o" // 成绩更新提醒
	const toAppid = "wx575f5d87fb66e69a"                             // 鸡音盒
	const appid = "wx9ef62ba2e3525812"                               // 鸡你太美小助手

	_, client := s.GetApp(appid)
	m := &message.TemplateMessage{
		ToUser:     openid,
		TemplateID: templateID,
		Data: map[string]*message.TemplateDataItem{
			"first":    {Value: "小黑子同学，你订阅的坤曲更新啦！\n➡️点击卡片畅听最新🐔乐⬅️", Color: ""},
			"keyword1": {Value: "二年级五班", Color: ""},
			"keyword2": {Value: "鸡你太美", Color: ""},
			"keyword3": {Value: "音乐", Color: ""},
			"keyword4": {Value: "两年半", Color: ""},
			"keyword5": {Value: "你坤哥", Color: ""},
			"remark":   {Value: "缘分一道桥、Havana、fantastic baby、empty love、Bad Romance、找朋友、樱花树下的约定、夜曲、夜的第七章、学猫叫、小鲤鱼历险记、小酒窝、吻别、天空之城、探坤、甩葱歌、世界这么大还是遇见你、如果我是DJ、如果感到幸福你就拍拍手、七龙珠、你好李焕英、南非世界杯、明日香、马保国接化发、卡塔尔世界杯、假烟假酒假朋友、黑猫警长、功夫足球、疯丫头、大笑江湖、大悲咒", Color: ""},
		},
		MiniProgram: struct {
			AppID    string `json:"appid"`
			PagePath string `json:"pagepath"`
		}{toAppid, "pages/index/index"},
	}

	return client.GetTemplate().Send(m)
}

func (s *WechatOcService) GetApp(appid string) (*conf.Application_App, *officialaccount.OfficialAccount) {
	return s.uc.GetApp(appid)
}
