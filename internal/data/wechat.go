package data

import (
	"context"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	errors2 "github.com/pkg/errors"
	"gorm.io/gorm/clause"
)

type wechatRepo struct {
	data *Data
	log  *log.Helper
}

func NewWechatRepo(data *Data, logger log.Logger) biz.WechatRepo {
	return &wechatRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *wechatRepo) GetRedisClient() redis.UniversalClient {
	return r.data.rdb
}

func (r *wechatRepo) GetUser(ctx context.Context, appid string, openid string) (m *biz.WechatUser, err error) {
	err = r.data.db.WithContext(ctx).First(m, "appid=? and openid=?", appid, openid).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return
}

func (r *wechatRepo) UpsertUser(ctx context.Context, m *biz.WechatUser) error {
	err := r.data.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "appid"}, {Name: "openid"}},
		UpdateAll: true, // 主键冲突时, 更新除主键的所有字段
	}).Create(m).Error
	if err != nil {
		return errors2.WithStack(err)
	}

	return nil
}
