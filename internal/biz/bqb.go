package biz

import (
	"context"
	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	errors2 "github.com/pkg/errors"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type BiaoQingBaoRepo interface {
	GetIndex(context.Context, string) ([]*BiaoQingBaoIndex, error)
	GetIndexNum(context.Context, string, []string) ([]*BiaoQingBaoIndexNum, error)
	GetList(context.Context, string, string, uint64, uint64) ([]*BiaoQingBao, error)
}

type BiaoQingBaoUseCase struct {
	repo BiaoQingBaoRepo
	log  *log.Helper
}

func NewBiaoQingBaoUseCase(repo BiaoQingBaoRepo, logger log.Logger) *BiaoQingBaoUseCase {
	return &BiaoQingBaoUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BiaoQingBaoUseCase) GetList(ctx context.Context, appid, typ string, page, pageSize uint64) (
	reply *v1.GetBqbListReply,
	err error,
) {
	baos, err := uc.repo.GetList(ctx, appid, typ, page, pageSize)
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	reply = &v1.GetBqbListReply{Results: make([]*v1.GetBqbListReply_Info, len(baos))}
	for i, bao := range baos {
		reply.Results[i] = &v1.GetBqbListReply_Info{
			Id:        uint64(bao.ID),
			Type:      bao.Type,
			ImagePath: bao.ImagePath,
		}
	}

	return reply, nil
}

func (uc *BiaoQingBaoUseCase) GetIndex(ctx context.Context, appid string) (reply *v1.GetBqbIndexReply, err error) {
	baos, err := uc.repo.GetIndex(ctx, appid)
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	types := make([]string, len(baos))
	for i, bao := range baos {
		types[i] = bao.Type
	}

	indexNums, err := uc.repo.GetIndexNum(ctx, appid, types)
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	indexNumMap := make(map[string]uint64)
	for _, num := range indexNums {
		indexNumMap[num.Type] = num.Num
	}

	reply = &v1.GetBqbIndexReply{Results: make([]*v1.GetBqbIndexReply_Info, len(baos))}
	var tmp *BiaoQingBaoIndex
	for i := range reply.Results {
		tmp = baos[i]
		reply.Results[i] = &v1.GetBqbIndexReply_Info{
			Name:      tmp.Name,
			Type:      tmp.Type,
			ImagePath: tmp.ImagePath,
		}
		if num, found := indexNumMap[tmp.Type]; found {
			reply.Results[i].Num = num
		}
	}

	return reply, nil
}
