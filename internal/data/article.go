package data

import (
	"context"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	errors2 "github.com/pkg/errors"

	"github.com/go-kratos/kratos/v2/log"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *articleRepo) GetArticle(ctx context.Context, appid, code string) (
	res *biz.Article, err error,
) {
	err = r.data.db.WithContext(ctx).
		First(&res, "appid=? and code=? and deleted_at is null", appid, code).
		Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return res, nil
}

func (r *articleRepo) GetArticles(ctx context.Context, appid, code string, page, pageSize int) (
	res []*biz.Article, err error,
) {
	err = r.data.db.WithContext(ctx).Limit(pageSize).Offset(page*pageSize).Order("sort").
		Find(&res, "appid=? and code=? and deleted_at is null", appid, code).
		Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return res, nil
}
