package data

import (
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
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
