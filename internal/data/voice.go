package data

import (
	"context"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	errors2 "github.com/pkg/errors"

	"github.com/go-kratos/kratos/v2/log"
)

type voiceRepo struct {
	data *Data
	log  *log.Helper
}

func NewVoiceRepo(data *Data, logger log.Logger) biz.VoiceRepo {
	return &voiceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *voiceRepo) GetDefault(ctx context.Context, appid string) (
	res *biz.Voice, err error,
) {
	err = r.data.db.WithContext(ctx).
		Where(&biz.Voice{
			Appid:   appid,
			Default: true,
		}).
		First(&res).
		Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return res, nil
}

func (r *voiceRepo) GetList(ctx context.Context, appid string) (
	res []*biz.Voice, err error,
) {
	err = r.data.db.WithContext(ctx).
		Select("appid", "code", "name", "default", "type", "mp_appid", "sort", "share_title", "share_image_url").
		Order("sort desc").
		Find(&res, "appid=?", appid).
		Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return res, nil
}

func (r *voiceRepo) Add(ctx context.Context, voice *biz.Voice) (err error) {
	err = r.data.db.WithContext(ctx).
		Save(voice).
		Error
	if err != nil {
		return errors2.WithStack(err)
	}
	return nil
}
