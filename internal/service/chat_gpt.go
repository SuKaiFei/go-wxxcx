package service

import (
	"context"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"

	pb "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatGptService struct {
	pb.UnimplementedChatGptServer
	uc *biz.ChatGPTUseCase
}

func NewChatGptService(uc *biz.ChatGPTUseCase) *ChatGptService {
	return &ChatGptService{uc: uc}
}

func (s *ChatGptService) GetChatGptCompletions(ctx context.Context, req *pb.GetChatGptCompletionsRequest) (*pb.GetChatGptCompletionsReply, error) {
	completionText, err := s.uc.Completions(ctx, req.Appid, req.Openid, req.Content)
	if err != nil {
		return nil, err
	}
	resp := &pb.GetChatGptCompletionsReply{
		Result:    completionText,
		ImagePath: "https://mmbiz.qpic.cn/mmbiz_jpg/Mhr8pCDXpQqoWjx3avyHfIMn9OJc93Pobae5GkdGX5E93SOBrichibgCNKxEn3JyCqeBBNmYkA5SHricEtyfMhs8A/0?wx_fmt=jpeg",
	}
	return resp, nil
}

func (s *ChatGptService) Ping(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
