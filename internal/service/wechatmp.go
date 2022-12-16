package service

import (
	"context"
	pb "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/security"
)

type WechatMpService struct {
	pb.UnimplementedWechatMpServer
	uc *biz.WechatUseCase
}

func NewWechatMpService(uc *biz.WechatUseCase) *WechatMpService {
	return &WechatMpService{uc: uc}
}

func (s *WechatMpService) LoginWechatMp(ctx context.Context, req *pb.LoginWechatMpRequest) (*pb.LoginWechatMpReply, error) {
	client := s.uc.GetMpApp(req.GetAppid())
	authCli := client.GetAuth()
	session, err := authCli.Code2Session(req.GetCode())
	if err != nil {
		return nil, err
	}
	if session.ErrCode != 0 {
		return nil, errors.New(session.Error())
	}
	reply := &pb.LoginWechatMpReply{
		Openid:  session.OpenID,
		Unionid: session.UnionID,
	}
	return reply, nil
}

func (s *WechatMpService) SecurityCheckMsg(ctx context.Context, req *pb.SecurityCheckMsgRequest) (
	*pb.SecurityCheckMsgReply,
	error,
) {
	client := s.uc.GetMpApp(req.GetAppid())
	securityCli := client.GetSecurity()
	securityRequest := &security.MsgCheckRequest{
		Scene:   2,
		OpenID:  req.Openid,
		Content: req.Content,
	}
	secCheck, err := securityCli.MsgCheck(securityRequest)
	if err != nil {
		return nil, err
	}
	if secCheck.ErrCode != 0 {
		return nil, errors.New(secCheck.Error())
	}
	if secCheck.Result.Label != 100 {
		log.Warnw(
			"message", "SecurityCheckMsg Label!=100",
			"content", req.Content,
			"appid", req.Appid,
			"openid", req.Openid,
		)
	}
	reply := &pb.SecurityCheckMsgReply{
		Suggest: string(secCheck.Result.Suggest),
		Label:   uint32(secCheck.Result.Label),
	}
	return reply, nil
}

func (s *WechatMpService) GetApp(appid string) *miniprogram.MiniProgram {
	return s.uc.GetMpApp(appid)
}
