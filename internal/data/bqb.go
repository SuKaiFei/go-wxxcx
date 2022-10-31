package data

import (
	"context"
	errors2 "github.com/pkg/errors"

	"github.com/SuKaiFei/go-wxxcx/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type bqbRepo struct {
	data *Data
	log  *log.Helper
}

func NewBqbRepo(data *Data, logger log.Logger) biz.BiaoQingBaoRepo {
	return &bqbRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *bqbRepo) GetIndex(ctx context.Context, appid string) (res []*biz.BiaoQingBaoIndex, err error) {
	err = r.data.db.WithContext(ctx).Find(&res, "appid=?", appid).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return res, nil
}

func (r *bqbRepo) GetIndexNum(ctx context.Context, appid string, types []string) (
	res []*biz.BiaoQingBaoIndexNum, err error,
) {
	err = r.data.db.WithContext(ctx).Table(new(biz.BiaoQingBao).TableName()).
		Select("COUNT( id ) num", "type").
		Group("type").
		Find(&res, "appid=? and type in ?", appid, types).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	return res, nil
}

func (r *bqbRepo) GetList(ctx context.Context, appid, typ string, page, pageSize uint64) (
	res []*biz.BiaoQingBao, err error,
) {
	err = r.data.db.WithContext(ctx).Limit(int(pageSize)).Offset(int(page*pageSize)).
		Order("updated_at desc").
		Find(&res, "appid=? and type=?", appid, typ).
		Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return res, nil
}
