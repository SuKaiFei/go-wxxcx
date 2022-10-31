package service

import (
	"context"
	"encoding/json"

	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"

	errors2 "github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

type VoiceService struct {
	v1.UnimplementedVoiceServer

	vuc *biz.VoiceUseCase
}

func NewVoiceService(uc *biz.VoiceUseCase) *VoiceService {
	return &VoiceService{vuc: uc}
}

func (s *VoiceService) GetVoiceList(ctx context.Context, in *v1.GetVoiceListRequest) (*v1.GetVoiceListReply, error) {
	if in.GetPageSize() > 50 {
		in.PageSize = 50
	}
	if in.GetPageSize() < 1 {
		in.PageSize = 10
	}
	voices, err := s.vuc.GetList(ctx, in.GetAppid())
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	reply := &v1.GetVoiceListReply{Results: make([]*v1.GetVoiceListReply_Info, len(voices))}
	for i, voice := range voices {
		reply.Results[i] = &v1.GetVoiceListReply_Info{
			Id:        uint64(voice.ID),
			Type:      "",
			ImagePath: "",
		}
	}

	return reply, nil
}

func (s *VoiceService) GetVoiceDefault(ctx context.Context, in *v1.GetVoiceDefaultRequest) (*v1.GetVoiceDefaultReply, error) {
	voice, err := s.vuc.GetDefault(ctx, in.GetAppid())
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	return voiceToPB(voice), nil
}

func (s *VoiceService) Ping(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {

	return &emptypb.Empty{}, nil
}

func voiceToPB(voice *biz.Voice) *v1.GetVoiceDefaultReply {
	var works string
	if len(voice.Works) > 0 {
		bytes, _ := json.Marshal(voice.Works)
		works = string(bytes)
	}
	return &v1.GetVoiceDefaultReply{
		Name:          voice.Name,
		Code:          voice.Code,
		Type:          uint32(voice.Type),
		MpAppid:       voice.MpAppid,
		Works:         works,
		ShareTitle:    voice.ShareTitle,
		ShareImageUrl: voice.ShareImageUrl,
	}
}
