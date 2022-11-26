package service

import (
	"context"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	errors2 "github.com/pkg/errors"

	pb "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MusicService struct {
	pb.UnimplementedMusicServer
	uc *biz.MusicUseCase
}

func NewMusicService(uc *biz.MusicUseCase) *MusicService {
	return &MusicService{uc: uc}
}

func (s *MusicService) GetMusicList(ctx context.Context, req *pb.GetMusicListRequest) (*pb.GetMusicListReply, error) {
	reply, err := s.uc.GetMusics(ctx, req.GetAppid(), req.GetCode(), req.GetPage(), req.GetPageSize())
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return reply, nil
}

func (s *MusicService) Ping(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
