package service

import (
	"context"
	pb "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatGptService struct {
	pb.UnimplementedChatGptServer
	uc *biz.ChatGPTUseCase
}

func NewChatGptService(uc *biz.ChatGPTUseCase) *ChatGptService {
	return &ChatGptService{uc: uc}
}

func (s *ChatGptService) GetChatGptHistory(ctx context.Context, req *pb.GetChatGptHistoryRequest) (*pb.GetChatGptHistoryReply, error) {
	res, err := s.uc.GetChatGptHistory(ctx, req.Openid, req.Id)
	if err != nil {
		return nil, err
	}
	reply := &pb.GetChatGptHistoryReply{
		Prompt: res.Prompt,
		Result: res.Result,
	}
	return reply, nil
}

func (s *ChatGptService) GetChatGptCompletions(ctx context.Context, req *pb.GetChatGptCompletionsRequest) (*pb.GetChatGptCompletionsReply, error) {
	reply, err := s.uc.Completions(ctx, req.Appid, req.Openid, req.Content)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *ChatGptService) GetAvailableCount(ctx context.Context, req *pb.GetAvailableCountRequest) (*pb.GetAvailableCountReply, error) {
	count, err := s.uc.GetTodaySendCount(ctx, req.Openid)
	if err != nil {
		return nil, err
	}
	resp := &pb.GetAvailableCountReply{
		Count: count,
	}
	return resp, nil
}

func (s *ChatGptService) CompleteAd(ctx context.Context, req *pb.CompleteAdRequest) (*emptypb.Empty, error) {
	err := s.uc.AddQuotaUnusedCount(ctx, req.Openid)
	if err != nil {
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (s *ChatGptService) Ping(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
