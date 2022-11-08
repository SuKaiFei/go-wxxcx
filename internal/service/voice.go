package service

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"

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
	voices, err := s.vuc.GetList(ctx, in.GetAppid())
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	reply := &v1.GetVoiceListReply{Results: make([]*v1.GetVoiceListReply_Info, len(voices))}
	for i, voice := range voices {
		reply.Results[i] = &v1.GetVoiceListReply_Info{
			Id:      uint64(voice.ID),
			Type:    uint32(voice.Type),
			Name:    voice.Name,
			Code:    voice.Code,
			MpAppid: voice.MpAppid,
		}
	}

	return reply, nil
}

func (s *VoiceService) GetVoiceDefault(ctx context.Context, in *v1.GetVoiceDefaultRequest) (*v1.GetVoiceReply, error) {
	voice, err := s.vuc.GetDefault(ctx, in.GetAppid())
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	return voiceToPB(voice), nil
}

func (s *VoiceService) GetVoiceById(ctx context.Context, in *v1.GetVoiceByIdRequest) (*v1.GetVoiceReply, error) {
	if in.GetId() == 0 {
		return s.GetVoiceDefault(ctx, &v1.GetVoiceDefaultRequest{Appid: in.GetAppid()})
	}
	voice, err := s.vuc.GetVoiceByID(ctx, in.GetId())
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	if voice.Appid != in.GetAppid() {
		return nil, errors.BadRequest("请求参数错误", "")
	}

	return voiceToPB(voice), nil
}

func (s *VoiceService) Ping(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {

	return &emptypb.Empty{}, nil
}

func voiceToPB(voice *biz.Voice) *v1.GetVoiceReply {
	var works string
	if len(voice.Works) > 0 {
		bytes, _ := json.Marshal(voice.Works)
		works = string(bytes)
	}
	return &v1.GetVoiceReply{
		Id:      uint64(voice.ID),
		Name:    voice.Name,
		Code:    voice.Code,
		Type:    uint32(voice.Type),
		MpAppid: voice.MpAppid,
		Works:   works,
		Share:   &v1.Share{Title: voice.ShareTitle, ImageUrl: voice.ShareImageUrl},
	}
}
