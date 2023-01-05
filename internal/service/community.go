package service

import (
	"context"
	pb "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	errors2 "github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type CommunityService struct {
	pb.UnimplementedCommunityServer
	uc       *biz.CommunityUseCase
	cosUc    *biz.CosUseCase
	wechatUc *biz.WechatUseCase
}

func NewCommunityService(uc *biz.CommunityUseCase, cosUc *biz.CosUseCase, wechatUc *biz.WechatUseCase) *CommunityService {
	return &CommunityService{uc: uc, cosUc: cosUc, wechatUc: wechatUc}
}

func (s *CommunityService) GetCommunitySettingNotice(ctx context.Context, req *pb.CommonRequest) (*pb.GetCommunitySettingNoticeReply, error) {
	m, err := s.uc.GetSettingNotice(ctx, req.Openid)
	if err != nil {
		return nil, err
	}

	return &pb.GetCommunitySettingNoticeReply{
		Id:                 uint64(m.ID),
		IsOpenLikeWork:     *m.IsOpenLikeWork,
		IsOpenLikeComment:  *m.IsOpenLikeComment,
		IsOpenCommentReply: *m.IsOpenCommentReply,
	}, nil
}

func (s *CommunityService) UpdateCommunitySettingNotice(ctx context.Context, req *pb.UpdateCommunitySettingNoticeRequest) (*emptypb.Empty, error) {
	err := s.uc.UpdateSettingNotice(ctx, req)
	if err != nil {
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (s *CommunityService) DeleteCommunityMyArticle(ctx context.Context, req *pb.DeleteCommunityByIdRequest) (*emptypb.Empty, error) {
	err := s.uc.DeleteMyArticle(ctx, req.Openid, req.Id)
	if err != nil {
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (s *CommunityService) DeleteCommunityMyComment(ctx context.Context, req *pb.DeleteCommunityByIdRequest) (*emptypb.Empty, error) {
	err := s.uc.DeleteMyComment(ctx, req.Openid, req.ArticleId, req.Id)
	if err != nil {
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (s *CommunityService) AddCommunityFeedback(ctx context.Context, req *pb.AddCommunityFeedbackRequest) (*emptypb.Empty, error) {
	_, err := s.uc.AddFeedback(ctx, req)
	if err != nil {
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (s *CommunityService) GetCommunityMyProfile(ctx context.Context, req *pb.GetCommunityMyProfileRequest) (*pb.GetCommunityMyProfileReply, error) {
	res, err := s.uc.GetMyProfile(ctx, req.GetOpenid())
	if err != nil {
		return nil, err
	}
	return &pb.GetCommunityMyProfileReply{
		Id:        uint64(res.ID),
		Username:  res.Username,
		AvatarUrl: res.Avatar,
	}, nil
}

func (s *CommunityService) UpdateCommunityMyProfile(ctx context.Context, req *pb.UpdateCommunityMyProfileRequest) (*emptypb.Empty, error) {
	check, err := s.wechatUc.MsgCheck(req.Appid, req.Openid, req.Username)
	if err != nil {
		return nil, err
	}
	if check.Result.Label != 100 {
		return nil, errors2.New("有违法违规文本内容,请重新输入")
	}

	err = s.uc.UpdateMyProfile(ctx, req)
	if err != nil {
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (s *CommunityService) GetCommunityMyArticleList(ctx context.Context, req *pb.GetCommunityArticleListRequest) (*pb.GetCommunityArticleListReply, error) {
	if req.GetPageSize() > 50 {
		req.PageSize = 50
	}
	req.PageSize = 5
	if req.GetPage() < 1 {
		req.Page = 0
	}
	reply, err := s.uc.GetMyArticleList(ctx, req.GetOpenid(), req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *CommunityService) GetCommunityCommentList(ctx context.Context, req *pb.GetCommunityCommentListRequest) (*pb.GetCommunityCommentListReply, error) {
	if req.GetPageSize() > 50 {
		req.PageSize = 50
	}
	if req.GetPage() < 1 {
		req.Page = 0
	}
	reply, err := s.uc.GetCommentList(ctx, req.Openid, req.ArticleId, req.CommentId, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *CommunityService) AddCommunityComment(ctx context.Context, req *pb.AddCommunityCommentRequest) (*pb.AddCommunityCommentReply, error) {
	check, err := s.wechatUc.MsgCheck(req.Appid, req.Openid, req.Content)
	if err != nil {
		return nil, err
	}
	if check.Result.Label != 100 {
		return nil, errors2.New("有违法违规文本内容,请重新输入")
	}

	reply, err := s.uc.AddComment(ctx, req)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *CommunityService) PushCommunityArticle(ctx context.Context, req *pb.PushCommunityArticleRequest) (*emptypb.Empty, error) {
	if len(req.Photos)+len(req.Content) == 0 {
		return nil, errors2.New("写点东西吧")
	}

	check, err := s.wechatUc.MsgCheck(req.Appid, req.Openid, req.Content)
	if err != nil {
		return nil, err
	}
	if check.Result.Label != 100 {
		return nil, errors2.New("有违法违规文本内容,请重新输入")
	}

	if len(req.Photos) > 0 {
		images := make([]string, len(req.Photos))
		for i, photo := range req.Photos {
			images[i] = photo.Url
		}

		//qrReader := qrcode.NewQRCodeReader()
		//for _, imageUrl := range images {
		// if strings.ToLower(path.Ext(imageUrl)) == ".gif" {
		// 	return nil, errors2.New("暂不支持GIF动态图上传")
		// }
		//	_, body, err := fasthttp.Get(nil, imageUrl)
		//	if err != nil {
		//		return nil, err
		//	}
		//	img, _, err := image.Decode(bytes.NewReader(body))
		//	if err != nil {
		//		if err == image.ErrFormat {
		//			continue
		//		}
		//		return nil, err
		//	}
		//
		//	// prepare BinaryBitmap
		//	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
		//	if err != nil {
		//		return nil, err
		//	}
		//
		//	// decode image
		//	result, _ := qrReader.Decode(bmp, nil)
		//	if result != nil && len(result.GetText()) > 0 {
		//		return nil, errors2.New("请删除包含二维码的图片")
		//	}
		//}

		auditing, err := s.cosUc.BatchImageAuditing(ctx, images)
		if err != nil {
			return nil, err
		}
		if !auditing {
			return nil, errors2.New("有违法违规图片,请重新选择")
		}
	}

	err = s.uc.CreateArticle(ctx, req)
	if err != nil {
		return nil, err
	}

	return new(emptypb.Empty), nil
}

func (s *CommunityService) UpdateCommunityLike(ctx context.Context, req *pb.UpdateCommunityLikeRequest) (*pb.UpdateCommunityLikeReply, error) {
	reply, err := s.uc.UpdateLike(ctx, req.Openid, req.Tid, biz.CommunityLikeType(req.Type))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *CommunityService) GetCommunityArticle(ctx context.Context, req *pb.GetCommunityArticleRequest) (*pb.GetCommunityArticleReply, error) {
	reply, err := s.uc.GetArticle(ctx, req.GetOpenid(), req.Id)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *CommunityService) GetCommunityArticleList(ctx context.Context, req *pb.GetCommunityArticleListRequest) (*pb.GetCommunityArticleListReply, error) {
	if req.GetPageSize() > 50 {
		req.PageSize = 50
	}
	req.PageSize = 5
	if req.GetPage() < 1 {
		req.Page = 0
	}
	reply, err := s.uc.GetArticleList(ctx, req.GetOpenid(), req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *CommunityService) GetCosCredential(ctx context.Context, req *pb.GetCosCredentialRequest) (*pb.GetCosCredentialReply, error) {
	res, err := s.cosUc.GetTempCredentials()
	if err != nil {
		return nil, err
	}

	p := &pb.GetCosCredentialReply{
		TmpSecretId:  res.Credentials.TmpSecretID,
		TmpSecretKey: res.Credentials.TmpSecretKey,
		SessionToken: res.Credentials.SessionToken,
		StartTime:    uint64(res.StartTime),
		ExpiredTime:  uint64(res.ExpiredTime),
	}
	return p, nil
}
