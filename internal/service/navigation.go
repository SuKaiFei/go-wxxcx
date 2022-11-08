package service

import (
	"context"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	errors2 "github.com/pkg/errors"

	pb "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
)

type NavigationService struct {
	pb.UnimplementedNavigationServer
	uc *biz.NavigationUseCase
}

func NewNavigationService(uc *biz.NavigationUseCase) *NavigationService {
	return &NavigationService{uc: uc}
}

func (s *NavigationService) GetNavigations(ctx context.Context, req *pb.GetNavigationsRequest) (*pb.GetNavigationsReply, error) {
	reply, err := s.uc.GetNavigations(ctx, req.GetAppid(), req.GetCode())
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return reply, nil
}
