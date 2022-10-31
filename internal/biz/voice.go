package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	errors2 "github.com/pkg/errors"
)

type VoiceRepo interface {
	GetList(context.Context, string) ([]*Voice, error)
	GetDefault(context.Context, string) (*Voice, error)
	Add(context.Context, *Voice) error
}

type VoiceUseCase struct {
	repo VoiceRepo
	log  *log.Helper
}

func NewVoiceUseCase(repo VoiceRepo, logger log.Logger) *VoiceUseCase {
	return &VoiceUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *VoiceUseCase) GetDefault(ctx context.Context, appid string) (
	reply *Voice,
	err error,
) {
	reply, err = uc.repo.GetDefault(ctx, appid)
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	return reply, nil
}

func (uc *VoiceUseCase) GetList(ctx context.Context, appid string) (
	reply []*Voice,
	err error,
) {
	reply, err = uc.repo.GetList(ctx, appid)
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	return reply, nil
}

func (uc *VoiceUseCase) Add(ctx context.Context, voice *Voice) (err error) {
	err = uc.repo.Add(ctx, voice)
	if err != nil {
		return errors2.WithStack(err)
	}

	return nil
}
