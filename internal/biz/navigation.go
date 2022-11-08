package biz

import (
	"context"
	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/go-kratos/kratos/v2/log"
	errors2 "github.com/pkg/errors"
)

type NavigationRepo interface {
	GetNavigations(context.Context, string, string) ([]*Navigation, error)
}

type NavigationUseCase struct {
	repo NavigationRepo
	log  *log.Helper
}

func NewNavigationUseCase(repo NavigationRepo, logger log.Logger) *NavigationUseCase {
	return &NavigationUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *NavigationUseCase) GetNavigations(ctx context.Context, appid, code string) (
	reply *v1.GetNavigationsReply,
	err error,
) {
	navigations, err := uc.repo.GetNavigations(ctx, appid, code)
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	reply = &v1.GetNavigationsReply{Results: make([]*v1.GetNavigationReply, len(navigations))}
	for i, navigation := range navigations {
		reply.Results[i] = &v1.GetNavigationReply{
			Type:      uint32(navigation.Type),
			ImagePath: navigation.ImagePath,
			Title:     navigation.Title,
			Describe:  navigation.Describe,
			Sort:      int64(navigation.Sort),
			MpAppid:   navigation.MpAppid,
			Url:       navigation.Url,
		}
	}

	return reply, nil
}
