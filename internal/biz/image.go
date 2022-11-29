package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type ImageUseCase struct {
	log *log.Helper
}

func NewImageUseCase(logger log.Logger) *ImageUseCase {
	return &ImageUseCase{log: log.NewHelper(logger)}
}
