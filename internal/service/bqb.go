package service

import (
	"context"
	errors2 "github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
)

type BqbService struct {
	v1.UnimplementedBqbServer

	uc *biz.BiaoQingBaoUseCase
}

func NewBqbService(uc *biz.BiaoQingBaoUseCase) *BqbService {
	return &BqbService{uc: uc}
}

func (s *BqbService) GetBqbIndex(ctx context.Context, in *v1.GetBqbIndexRequest) (*v1.GetBqbIndexReply, error) {
	reply, err := s.uc.GetIndex(ctx, in.GetAppid())
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return reply, nil
}

func (s *BqbService) GetBqbList(ctx context.Context, in *v1.GetBqbListRequest) (*v1.GetBqbListReply, error) {
	if in.GetPageSize() > 50 {
		in.PageSize = 50
	}
	if in.GetPageSize() < 1 {
		in.PageSize = 10
	}
	reply, err := s.uc.GetList(ctx, in.GetAppid(), in.GetType(), in.GetPage(), in.GetPageSize())
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return reply, nil
}

func (s *BqbService) Ping(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
