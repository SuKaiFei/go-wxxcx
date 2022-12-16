package data

import (
	"context"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	errors2 "github.com/pkg/errors"

	"github.com/go-kratos/kratos/v2/log"
)

type ChatGPTRepo struct {
	data *Data
	log  *log.Helper
}

func NewChatGPTRepo(data *Data, logger log.Logger) biz.ChatGPTRepo {
	return &ChatGPTRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ChatGPTRepo) Get(ctx context.Context, code string, prompt string) (m *biz.ChatGPT, err error) {
	err = r.data.db.WithContext(ctx).First(&m, "code=? and prompt=?", code, prompt).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return
}

func (r *ChatGPTRepo) Add(ctx context.Context, m *biz.ChatGPT) error {
	return r.data.db.WithContext(ctx).Create(m).Error
}

func (r *ChatGPTRepo) CountByOpenid(ctx context.Context, openid string) (m int64, err error) {
	err = r.data.db.WithContext(ctx).Model(new(biz.ChatGPT)).Where("openid=?", openid).
		Where("created_at< TIMESTAMPADD(MICROSECOND,-1,DATE_ADD(CURDATE(),INTERVAL 1 DAY)) AND created_at> TIMESTAMP (CURDATE())").
		Count(&m).Error
	if err != nil {
		return 0, errors2.WithStack(err)
	}
	return
}

func (r *ChatGPTRepo) GetList(ctx context.Context, appid string) (
	res []*biz.ChatGPT, err error,
) {
	err = r.data.db.WithContext(ctx).
		Order("sort").
		Find(&res, "appid=?", appid).
		Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return res, nil
}
