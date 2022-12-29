package data

import (
	"context"
	"fmt"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	errors2 "github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
	"time"

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

func (r *ChatGPTRepo) Get(ctx context.Context, id uint) (m *biz.ChatGPT, err error) {
	err = r.data.db.WithContext(ctx).First(&m, "id = ?", id).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return
}

func (r *ChatGPTRepo) Add(ctx context.Context, m *biz.ChatGPT) (uint, error) {
	err := r.data.db.WithContext(ctx).Create(m).Error
	if err != nil {
		return 0, errors2.WithStack(err)
	}
	return m.ID, nil
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

func (r *ChatGPTRepo) Search(ctx context.Context, code string, prompt string) (
	res *biz.ChatGPT, err error,
) {
	err = r.data.db.WithContext(ctx).
		First(&res, "type=1 and code=? and prompt=?", code, prompt).
		Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return
}

func (r *ChatGPTRepo) GetTodayQuota(ctx context.Context, openid string) (m *biz.ChatGPTQuota, err error) {
	now := time.Now()
	date, _ := strconv.Atoi(fmt.Sprintf("%0.4d%0.2d%0.2d", now.Year(), now.Month(), now.Day()))
	err = r.data.db.WithContext(ctx).
		Where("openid=? and date=?", openid, date).
		First(&m).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			m = &biz.ChatGPTQuota{
				Openid:      openid,
				Date:        uint64(date),
				UnusedCount: 1,
			}
			err = r.data.db.WithContext(ctx).Create(m).Error
			if err == nil {
				return m, nil
			}
		}
		return nil, errors2.WithStack(err)
	}
	return
}

func (r *ChatGPTRepo) AddQuotaUseCount(ctx context.Context, openid string) error {
	now := time.Now()
	date, _ := strconv.Atoi(fmt.Sprintf("%0.4d%0.2d%0.2d", now.Year(), now.Month(), now.Day()))

	err := r.data.db.WithContext(ctx).Model(new(biz.ChatGPTQuota)).
		Where("date=? and openid=?", date, openid).
		Update("use_count", gorm.Expr("use_count + 1")).Error
	if err != nil {
		return errors2.WithStack(err)
	}

	return nil
}

func (r *ChatGPTRepo) AddQuotaUnusedCount(ctx context.Context, openid string) error {
	now := time.Now()
	date, _ := strconv.Atoi(fmt.Sprintf("%0.4d%0.2d%0.2d", now.Year(), now.Month(), now.Day()))

	err := r.data.db.WithContext(ctx).Model(new(biz.ChatGPTQuota)).
		Where("date=? and openid=?", date, openid).
		Update("unused_count", gorm.Expr("unused_count + 1")).Error
	if err != nil {
		return errors2.WithStack(err)
	}

	return nil
}
