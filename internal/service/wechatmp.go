package service

import (
	"context"
	pb "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/medivhzhan/weapp/v3"
	"github.com/medivhzhan/weapp/v3/auth"
)

type WechatMpService struct {
	pb.UnimplementedWechatMpServer
	uc *biz.WechatMpUseCase
}

func NewWechatMpService(uc *biz.WechatMpUseCase) *WechatMpService {
	return &WechatMpService{uc: uc}
}

func (s *WechatMpService) LoginWechatMp(ctx context.Context, req *pb.LoginWechatMpRequest) (*pb.LoginWechatMpReply, error) {
	confApp, client := s.GetApp(req.GetAppid())
	authCli := client.NewAuth()
	sessionRequest := &auth.Code2SessionRequest{
		Appid:     req.GetAppid(),
		Secret:    confApp.GetSecret(),
		JsCode:    req.GetCode(),
		GrantType: "authorization_code",
	}
	session, err := authCli.Code2Session(sessionRequest)
	if err != nil {
		return nil, err
	}
	reply := &pb.LoginWechatMpReply{
		Openid:     session.Openid,
		SessionKey: session.SessionKey,
		Unionid:    session.Unionid,
	}
	return reply, nil
}

func (s *WechatMpService) GetApp(appid string) (*conf.Application_App, *weapp.Client) {
	return s.uc.GetApp(appid)
}
