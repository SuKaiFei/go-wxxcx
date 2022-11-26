package biz

import (
	"context"
	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/go-kratos/kratos/v2/log"
	errors2 "github.com/pkg/errors"
)

type MusicRepo interface {
	GetMusics(context.Context, string, string, int, int) ([]*Music, error)
}

type MusicUseCase struct {
	repo MusicRepo
	log  *log.Helper
}

func NewMusicUseCase(repo MusicRepo, logger log.Logger) *MusicUseCase {
	return &MusicUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *MusicUseCase) GetMusics(ctx context.Context, appid, code string, page, pageSize uint64) (
	reply *v1.GetMusicListReply,
	err error,
) {
	musics, err := uc.repo.GetMusics(ctx, appid, code, int(page), int(pageSize))
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	reply = &v1.GetMusicListReply{Results: make([]*v1.GetMusicListReply_Info, len(musics))}
	for i, music := range musics {
		reply.Results[i] = &v1.GetMusicListReply_Info{
			Id:        uint64(music.ID),
			Name:      music.Name,
			Duration:  music.Duration,
			Singer:    music.Singer,
			Url:       music.Url,
			ImagePath: music.ImagePath,
			Share: &v1.Share{
				Title:    music.ShareTitle,
				ImageUrl: music.ShareImageUrl,
			},
		}
	}

	return reply, nil
}
