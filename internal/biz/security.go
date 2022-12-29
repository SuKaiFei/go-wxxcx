package biz

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	errors2 "github.com/pkg/errors"
)

type SecurityRepo interface {
	AddSign(context.Context, string, string, string, string) error
}

type SecurityUseCase struct {
	repo SecurityRepo
	log  *log.Helper
}

func NewSecurityUseCase(repo SecurityRepo, logger log.Logger) *SecurityUseCase {
	return &SecurityUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SecurityUseCase) VerifySign(ctx context.Context, reqUrl, appid, openid, sign string) (err error) {
	md := md5.New()
	md.Write([]byte(reqUrl))
	reqFlag := fmt.Sprintf("%+x", md.Sum(nil))
	err = uc.repo.AddSign(ctx, reqFlag, appid, openid, sign)
	if err != nil {
		return errors2.WithStack(err)
	}
	return nil
}
