package data

import (
	"context"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	errors2 "github.com/pkg/errors"

	"github.com/go-kratos/kratos/v2/log"
)

type navigationRepo struct {
	data *Data
	log  *log.Helper
}

func NewNavigationRepo(data *Data, logger log.Logger) biz.NavigationRepo {
	return &navigationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *navigationRepo) GetNavigations(ctx context.Context, appid, code string) (
	res []*biz.Navigation, err error,
) {
	err = r.data.db.WithContext(ctx).Order("sort").
		Find(&res, "appid=? and code=? and deleted_at is null", appid, code).
		Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return res, nil
}
