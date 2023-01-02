package service

import (
	"context"
	pb "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/silenceper/wechat/v2/miniprogram"
)

type WechatMpService struct {
	pb.UnimplementedWechatMpServer
	uc          *biz.WechatUseCase
	communityUC *biz.CommunityUseCase
}

func NewWechatMpService(uc *biz.WechatUseCase, communityUC *biz.CommunityUseCase) *WechatMpService {
	return &WechatMpService{uc: uc, communityUC: communityUC}
}

func (s *WechatMpService) LoginWechatMp(ctx context.Context, req *pb.LoginWechatMpRequest) (*pb.LoginWechatMpReply, error) {
	session, err := s.uc.Code2Session(req.Appid, req.Code)
	if err != nil {
		return nil, err
	}
	if req.Appid == "wxec615f70feb4e93c" {
		go func() {
			err2 := s.communityUC.UpdateUserUnionid(context.Background(), session.OpenID, session.UnionID)
			log.Errorw("msg", "communityUC.UpdateUserUnionid", "err", err2)
		}()
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
