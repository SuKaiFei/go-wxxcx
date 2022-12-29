package service

import (
	"context"
	pb "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	"github.com/pkg/errors"
	"github.com/silenceper/wechat/v2/miniprogram"
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
	secCheck, err := s.uc.MsgCheck(req.Appid, req.Openid, req.Content)
	if err != nil {
		return nil, err
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
