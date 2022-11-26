package data

import (
	"context"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	errors2 "github.com/pkg/errors"

	"github.com/go-kratos/kratos/v2/log"
)

type musicRepo struct {
	data *Data
	log  *log.Helper
}

func NewMusicRepo(data *Data, logger log.Logger) biz.MusicRepo {
	return &musicRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *musicRepo) GetMusics(ctx context.Context, appid, code string, page, pageSize int) (
	res []*biz.Music, err error,
) {
	err = r.data.db.WithContext(ctx).Limit(pageSize).Offset(page*pageSize).Order("sort").
		Find(&res, "appid=? and code=?", appid, code).
		Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return res, nil
}
