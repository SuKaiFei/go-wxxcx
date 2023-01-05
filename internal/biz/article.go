package biz

import (
	"context"
	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/go-kratos/kratos/v2/log"
	errors2 "github.com/pkg/errors"
)

type ArticleRepo interface {
	GetArticle(context.Context, string, string) (*Article, error)
	GetArticles(context.Context, string, string, int, int) ([]*Article, error)
}

type ArticleUseCase struct {
	repo ArticleRepo
	log  *log.Helper
}

func NewArticleUseCase(repo ArticleRepo, logger log.Logger) *ArticleUseCase {
	return &ArticleUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ArticleUseCase) GetArticle(ctx context.Context, appid, code string) (
	reply *v1.GetArticleReply,
	err error,
) {
	article, err := uc.repo.GetArticle(ctx, appid, code)
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	return &v1.GetArticleReply{
		ImagePath: article.ImagePath,
		Title:     article.Title,
		Content:   article.Content,
		Sort:      uint64(article.Sort),
	}, nil
}

func (uc *ArticleUseCase) GetArticles(ctx context.Context, appid, code string, page, pageSize uint64) (
	reply *v1.GetArticlesReply,
	err error,
) {
	articles, err := uc.repo.GetArticles(ctx, appid, code, int(page), int(pageSize))
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	reply = &v1.GetArticlesReply{Results: make([]*v1.GetArticleReply, len(articles))}
	for i, article := range articles {
		reply.Results[i] = &v1.GetArticleReply{
			ImagePath: article.ImagePath,
			Title:     article.Title,
			Content:   article.Content,
			Sort:      uint64(article.Sort),
		}
	}

	return reply, nil
}
