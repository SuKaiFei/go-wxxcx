package data

import (
	"context"
	"fmt"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	errors2 "github.com/pkg/errors"
	"time"
)

type securityRepo struct {
	data *Data
	log  *log.Helper
}

func NewSecurityRepo(data *Data, logger log.Logger) biz.SecurityRepo {
	return &securityRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *securityRepo) AddSign(ctx context.Context, reqFlag, appid, openid, sign string) error {
	conn := r.data.rdb.Conn(ctx)
	defer conn.Close()
	isOk, err := conn.SetNX(ctx, fmt.Sprintf("sign:%s:%s:%s:%s", appid, openid, reqFlag, sign), 1, 5*time.Minute).Result()
	if err != nil {
		return errors2.WithStack(err)
	}
	if !isOk {
		return errors2.New("sign existed")
	}
	return nil
}
