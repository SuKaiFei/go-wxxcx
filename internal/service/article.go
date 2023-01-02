package service

import (
	"context"

	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	errors2 "github.com/pkg/errors"

	pb "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
)

type ArticleService struct {
	pb.UnimplementedArticleServer

	uc *biz.ArticleUseCase
}

func NewArticleService(uc *biz.ArticleUseCase) *ArticleService {
	return &ArticleService{uc: uc}
}

func (s *ArticleService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	reply, err := s.uc.GetArticle(ctx, req.GetAppid(), req.GetCode())
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return reply, nil
}

func (s *ArticleService) GetArticles(ctx context.Context, req *pb.GetArticlesRequest) (*pb.GetArticlesReply, error) {
	if req.GetPageSize() > 50 {
		req.PageSize = 50
	}
	if req.GetPage() < 1 {
		req.Page = 0
	}
	reply, err := s.uc.GetArticles(ctx, req.GetAppid(), req.GetCode(), req.GetPage(), req.GetPageSize())
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return reply, nil
}
