package biz

import (
	"context"
	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/go-kratos/kratos/v2/log"
	errors2 "github.com/pkg/errors"
	"gorm.io/gorm"
)

var (
	defaultCommentUser = &CommunityUser{
		Username: "微信用户",
		Avatar:   "https://mmbiz.qpic.cn/mmbiz_png/Mhr8pCDXpQqoWjx3avyHfIMn9OJc93Po20icsy9C9Qsd8lu6N9OUpgNetmssiaA96WgibiaWKVxHRz0oIz78JOA2Nw/0?wx_fmt=png",
	}
)

type CommunityRepo interface {
	GetUser(context.Context, string) (*CommunityUser, error)
	AddUser(context.Context, *CommunityUser) (uint, error)
	UpdateUser(context.Context, uint, *CommunityUser) error
	GetUserListByOpenid(context.Context, []string) ([]*CommunityUser, error)
	GetComment(context.Context, uint) (m *CommunityComment, err error)
	GetCommentList(context.Context, uint, uint, int, int) (m []*CommunityComment, err error)
	GetArticleList(context.Context, int, int) (m []*CommunityArticle, err error)
	GetArticleListByOpenid(context.Context, string, int, int) (m []*CommunityArticle, err error)
	GetArticle(context.Context, uint) (*CommunityArticle, error)
	GetArticleLikeByOpenid(context.Context, string, []uint) (m []*CommunityLike, err error)
	GetCommentLikeByOpenid(context.Context, string, []uint) (m []*CommunityLike, err error)
	GetLike(context.Context, string, uint, CommunityLikeType) (*CommunityLike, error)
	AddLike(context.Context, *CommunityLike) (uint, error)
	AddArticle(context.Context, *CommunityArticle) (uint, error)
	AddComment(context.Context, *CommunityComment) (uint, error)
	DeleteArticle(context.Context, string, uint) (int64, error)
	DeleteComment(context.Context, string, uint) (int64, error)
	AddFeedback(context.Context, *CommunityFeedback) (uint, error)
	UpdateLikeStatus(context.Context, uint, CommunityLikeStatus) error
	UpdateArticleLikeCount(context.Context, uint, int) error
	UpdateArticleCommentCount(context.Context, uint, int) error
	UpdateCommentLikeCount(context.Context, uint, int) error
	UpdateCommentCount(context.Context, uint, int) error
}

type CommunityUseCase struct {
	repo CommunityRepo
	log  *log.Helper
}

func NewCommunityUseCase(repo CommunityRepo, logger log.Logger) *CommunityUseCase {
	return &CommunityUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *CommunityUseCase) AddComment(ctx context.Context, req *v1.AddCommunityCommentRequest) (*v1.AddCommunityCommentReply, error) {
	m := &CommunityComment{
		ArticleID:      req.ArticleId,
		Openid:         req.Openid,
		ReplyOpenid:    req.ReplyOpenid,
		TopReplyOpenid: req.TopReplyOpenid,
		Content:        req.Content,
		ReplyID:        req.ReplyId,
		TopReplyID:     req.TopReplyId,
	}
	_, err := uc.repo.AddComment(ctx, m)
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	err = uc.repo.UpdateArticleCommentCount(ctx, uint(req.ArticleId), 1)
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	if req.TopReplyId > 0 {
		err = uc.repo.UpdateCommentCount(ctx, uint(req.TopReplyId), 1)
		if err != nil {
			return nil, errors2.WithStack(err)
		}
	}

	reply, err := uc.fillCommentReply(ctx, req.Openid, []*CommunityComment{m})
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	return reply[0], nil
}

func (uc *CommunityUseCase) DeleteMyArticle(ctx context.Context, openid string, id uint64) error {
	_, err := uc.repo.DeleteArticle(ctx, openid, uint(id))
	if err != nil {
		return errors2.WithStack(err)
	}

	return nil
}

func (uc *CommunityUseCase) DeleteMyComment(ctx context.Context, openid string, articleID, id uint64) error {
	deleteCount, err := uc.repo.DeleteComment(ctx, openid, uint(id))
	if err != nil {
		return errors2.WithStack(err)
	}
	if deleteCount > 0 {
		err = uc.repo.UpdateArticleCommentCount(ctx, uint(articleID), -1)
		if err != nil {
			return errors2.WithStack(err)
		}
	}

	return nil
}

func (uc *CommunityUseCase) AddFeedback(ctx context.Context, req *v1.AddCommunityFeedbackRequest) (uint, error) {
	m := &CommunityFeedback{
		Openid:  req.Openid,
		Content: req.Content,
		Type:    CommunityLikeType(req.Type),
		Tid:     req.Id,
	}
	_, err := uc.repo.AddFeedback(ctx, m)
	if err != nil {
		return 0, errors2.WithStack(err)
	}

	return m.ID, nil
}

func (uc *CommunityUseCase) CreateArticle(ctx context.Context, req *v1.PushCommunityArticleRequest) error {
	m := &CommunityArticle{
		Openid:  req.Openid,
		Content: req.Content,
		Photos:  req.Photos,
	}
	if _, err := uc.repo.AddArticle(ctx, m); err != nil {
		return errors2.WithStack(err)
	}

	return nil
}

func (uc *CommunityUseCase) GetUserMapByOpenid(ctx context.Context, openIDs []string) (map[string]*CommunityUser, error) {
	items, err := uc.repo.GetUserListByOpenid(ctx, openIDs)
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	itemMap := make(map[string]*CommunityUser, len(items))
	for _, item := range items {
		itemMap[item.Openid] = item
	}

	return itemMap, nil
}

func (uc *CommunityUseCase) GetArticle(ctx context.Context, openid string, id uint64) (
	*v1.GetCommunityArticleReply, error,
) {
	item, err := uc.repo.GetArticle(ctx, uint(id))
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	reply, err := uc.fillArticleReply(ctx, openid, append([]*CommunityArticle{}, item))
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	return reply[0], nil
}

func (uc *CommunityUseCase) GetMyProfile(ctx context.Context, openid string) (
	*CommunityUser, error,
) {
	item, err := uc.repo.GetUser(ctx, openid)
	if err != nil && errors2.Cause(err) != gorm.ErrRecordNotFound {
		return nil, errors2.WithStack(err)
	}
	if item == nil {
		item = &CommunityUser{
			Openid:   openid,
			Username: defaultCommentUser.Username,
			Avatar:   defaultCommentUser.Avatar,
		}
		_, err = uc.repo.AddUser(ctx, item)
		if err != nil {
			return nil, errors2.WithStack(err)
		}
	}

	return item, nil
}

func (uc *CommunityUseCase) UpdateMyProfile(ctx context.Context, req *v1.UpdateCommunityMyProfileRequest) error {
	m := &CommunityUser{
		Username: req.Username,
		Avatar:   req.AvatarUrl,
	}
	err := uc.repo.UpdateUser(ctx, uint(req.Id), m)
	if err != nil {
		return errors2.WithStack(err)
	}

	return nil
}

func (uc *CommunityUseCase) GetArticleList(ctx context.Context, openid string, page, pageSize uint64) (
	*v1.GetCommunityArticleListReply, error,
) {
	items, err := uc.repo.GetArticleList(ctx, int(page), int(pageSize))
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	reply, err := uc.fillArticleReply(ctx, openid, items)
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	return &v1.GetCommunityArticleListReply{Results: reply}, nil
}

func (uc *CommunityUseCase) GetCommentList(ctx context.Context, openid string, aID, comID, page, pageSize uint64) (
	*v1.GetCommunityCommentListReply, error,
) {
	items, err := uc.repo.GetCommentList(ctx, uint(aID), uint(comID), int(page), int(pageSize))
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	reply, err := uc.fillCommentReply(ctx, openid, items)
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	return &v1.GetCommunityCommentListReply{Results: reply}, nil
}

func (uc *CommunityUseCase) GetMyArticleList(ctx context.Context, openid string, page, pageSize uint64) (
	*v1.GetCommunityArticleListReply, error,
) {
	items, err := uc.repo.GetArticleListByOpenid(ctx, openid, int(page), int(pageSize))
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	reply, err := uc.fillArticleReply(ctx, openid, items)
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	return &v1.GetCommunityArticleListReply{Results: reply}, nil
}

func (uc *CommunityUseCase) fillCommentReply(ctx context.Context, openid string, items []*CommunityComment) (
	[]*v1.AddCommunityCommentReply, error,
) {
	ids := make([]uint, len(items))
	pushOpenIDs := make([]string, 0, len(items)*2)
	for i, item := range items {
		ids[i] = item.ID
		pushOpenIDs = append(pushOpenIDs, item.Openid)
		if item.ReplyID > 0 {
			pushOpenIDs = append(pushOpenIDs, item.ReplyOpenid)
		}
	}

	userMap, err := uc.GetUserMapByOpenid(ctx, pushOpenIDs)
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	likes, err := uc.repo.GetCommentLikeByOpenid(ctx, openid, ids)
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	likeMap := make(map[uint]struct{}, len(likes))
	for _, like := range likes {
		if like.Status != CommunityLikeStatusLike {
			continue
		}
		likeMap[uint(like.Tid)] = struct{}{}
	}

	res := make([]*v1.AddCommunityCommentReply, len(items))
	for i, item := range items {
		user, found := userMap[item.Openid]
		if !found {
			user = defaultCommentUser
		}
		replyUser, found := userMap[item.ReplyOpenid]
		if !found {
			replyUser = defaultCommentUser
		}
		_, isLike := likeMap[item.ID]

		if item.LikeCount < 0 {
			item.LikeCount = 0
		}
		if item.CommentCount < 0 {
			item.CommentCount = 0
		}

		res[i] = &v1.AddCommunityCommentReply{
			Id:             uint64(item.ID),
			Content:        item.Content,
			Userinfo:       &v1.CommunityUser{Id: uint64(user.ID), Openid: user.Openid, Username: user.Username, AvatarUrl: user.Avatar},
			TopReplyId:     item.TopReplyID,
			ReplyId:        item.ReplyID,
			TopReplyOpenid: item.TopReplyOpenid,
			LikeCount:      uint64(item.LikeCount),
			CommentCount:   uint64(item.CommentCount),
			IsLike:         isLike,
			CommentTime:    item.CreatedAt.Format("01-02 15:04"),
		}
		if item.ReplyID > 0 {
			res[i].ReplyUserinfo = &v1.CommunityUser{
				Id:        uint64(replyUser.ID),
				Openid:    replyUser.Openid,
				Username:  replyUser.Username,
				AvatarUrl: replyUser.Avatar,
			}
		}
	}

	return res, nil
}

func (uc *CommunityUseCase) fillArticleReply(ctx context.Context, openid string, items []*CommunityArticle) (
	[]*v1.GetCommunityArticleReply, error,
) {

	ids := make([]uint, len(items))
	pushOpenIDs := make([]string, len(items))
	for i, item := range items {
		ids[i] = item.ID
		pushOpenIDs[i] = item.Openid
	}

	userMap, err := uc.GetUserMapByOpenid(ctx, pushOpenIDs)
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	likes, err := uc.repo.GetArticleLikeByOpenid(ctx, openid, ids)
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	likeMap := make(map[uint]struct{}, len(likes))
	for _, like := range likes {
		if like.Status != CommunityLikeStatusLike {
			continue
		}
		likeMap[uint(like.Tid)] = struct{}{}
	}

	res := make([]*v1.GetCommunityArticleReply, len(items))
	for i, item := range items {
		user, found := userMap[item.Openid]
		if !found {
			user = defaultCommentUser
		}
		_, isLike := likeMap[item.ID]
		if !found {
			user = defaultCommentUser
		}

		if item.LikeCount < 0 {
			item.LikeCount = 0
		}
		if item.CommentCount < 0 {
			item.CommentCount = 0
		}

		res[i] = &v1.GetCommunityArticleReply{
			Id:            uint64(item.ID),
			PubUserName:   user.Username,
			PubUserAvatar: user.Avatar,
			PubContent:    item.Content,
			PubTime:       item.CreatedAt.Format("01-02 15:04"),
			LikeCount:     uint64(item.LikeCount),
			ComCount:      uint64(item.CommentCount),
			IsLike:        isLike,
			Photos:        item.Photos,
			Openid:        item.Openid,
		}
	}

	return res, nil
}

func (uc *CommunityUseCase) UpdateLike(ctx context.Context, openid string, tid uint64, typ CommunityLikeType) (
	*v1.UpdateCommunityLikeReply, error,
) {
	like, err := uc.repo.GetLike(ctx, openid, uint(tid), typ)
	if err != nil && errors2.Cause(err) != gorm.ErrRecordNotFound {
		return nil, errors2.WithStack(err)
	}
	updateTo := CommunityLikeStatusLike
	if like == nil {
		m := &CommunityLike{
			Openid: openid,
			Type:   typ,
			Tid:    tid,
			Status: CommunityLikeStatusLike,
		}
		_, err = uc.repo.AddLike(ctx, m)
		if err != nil {
			return nil, errors2.WithStack(err)
		}
	} else {
		if like.Status != CommunityLikeStatusUnlike {
			updateTo = CommunityLikeStatusUnlike
		}
		if like.Status != CommunityLikeStatusUnlike {
			updateTo = CommunityLikeStatusUnlike
		}
		err = uc.repo.UpdateLikeStatus(ctx, like.ID, updateTo)
		if err != nil {
			return nil, errors2.WithStack(err)
		}
	}

	step := 1
	if updateTo != CommunityLikeStatusLike {
		step = step * -1
	}

	if typ == CommunityLikeTypeArticle {
		err = uc.repo.UpdateArticleLikeCount(ctx, uint(tid), step)
		if err != nil {
			return nil, errors2.WithStack(err)
		}
	} else if typ == CommunityLikeTypeComment {
		err = uc.repo.UpdateCommentLikeCount(ctx, uint(tid), step)
		if err != nil {
			return nil, errors2.WithStack(err)
		}
	}

	return &v1.UpdateCommunityLikeReply{IsLike: updateTo == CommunityLikeStatusLike}, nil
}
