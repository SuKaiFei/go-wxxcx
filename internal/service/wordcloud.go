package service

import (
	"context"
	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	errors2 "github.com/pkg/errors"
	"strings"
)

type WordcloudService struct {
	v1.UnimplementedWordcloudServer
	uc   *biz.WordcloudUseCase
	wcUC *biz.WechatUseCase
}

func NewWordcloudService(uc *biz.WordcloudUseCase, wcUC *biz.WechatUseCase) *WordcloudService {
	return &WordcloudService{uc: uc, wcUC: wcUC}
}

func (s *WordcloudService) GenerateWordcloudImage(ctx context.Context, req *v1.GenerateWordcloudImageRequest) (
	reply *v1.UploadImageReply,
	err error,
) {
	check, err := s.wcUC.MsgCheck(req.Appid, req.Openid, strings.Join(req.Words, ","))
	if err != nil {
		return nil, err
	}
	if check.Result.Label != 100 {
		return nil, errors2.New("有违法违规文本内容,请重新输入")
	}

	return s.uc.GenerateImage(ctx, req.ImagePath, append(req.Words, "小程序表情制作小工具"))
}
