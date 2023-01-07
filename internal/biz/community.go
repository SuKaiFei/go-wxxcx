package biz

import (
	"context"
	"fmt"
	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	errors2 "github.com/pkg/errors"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"gorm.io/gorm"
	"time"
)

var (
	defaultSettingNotice = true
	defaultCommentUser   = &CommunityUser{
		Username:     "微信用户",
		Avatar:       "https://mmbiz.qpic.cn/mmbiz_png/Mhr8pCDXpQqoWjx3avyHfIMn9OJc93Po20icsy9C9Qsd8lu6N9OUpgNetmssiaA96WgibiaWKVxHRz0oIz78JOA2Nw/0?wx_fmt=png",
		Introduction: "这是一个什么也没有写的小黑子",
	}
)

type CommunityRepo interface {
	AddNoticeHistory(context.Context, *CommunityNoticeHistory) (uint, error)
	GetUser(context.Context, string) (*CommunityUser, error)
	AddUser(context.Context, *CommunityUser) (uint, error)
	UpdateUser(context.Context, uint, *CommunityUser) error
	GetSettingNotice(context.Context, string) (*CommunitySettingNotice, error)
	AddSettingNotice(context.Context, *CommunitySettingNotice) (uint, error)
	UpdateSettingNotice(context.Context, uint, *CommunitySettingNotice) error
	UpdateUserUnionid(context.Context, string, string) error
	GetUserListByOpenid(context.Context, []string) ([]*CommunityUser, error)
	GetComment(context.Context, uint) (m *CommunityComment, err error)
	GetCommentList(context.Context, uint, uint, int, int) (m []*CommunityComment, err error)
	GetArticleList(context.Context, int, int) (m []*CommunityArticle, err error)
	GetTopArticleList(context.Context) (m []*CommunityArticle, err error)
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
	repo     CommunityRepo
	cosUC    *CosUseCase
	wechatUc *WechatUseCase
	log      *log.Helper
}

func NewCommunityUseCase(repo CommunityRepo, cosUC *CosUseCase, wechatUc *WechatUseCase, logger log.Logger) *CommunityUseCase {
	return &CommunityUseCase{
		cosUC:    cosUC,
		wechatUc: wechatUc,
		repo:     repo,
		log:      log.NewHelper(logger),
	}
}

func (uc *CommunityUseCase) GetSettingNotice(ctx context.Context, openid string) (*CommunitySettingNotice, error) {
	m, err := uc.repo.GetSettingNotice(ctx, openid)
	if err != nil && errors2.Cause(err) != gorm.ErrRecordNotFound {
		return nil, errors2.WithStack(err)
	}
	if m == nil {
		user, err := uc.GetMyProfile(ctx, openid)
		if err != nil {
			return nil, errors2.WithStack(err)
		}
		m = &CommunitySettingNotice{
			Openid:             openid,
			Unionid:            user.Unionid,
			IsOpenLikeWork:     &defaultSettingNotice,
			IsOpenLikeComment:  &defaultSettingNotice,
			IsOpenCommentReply: &defaultSettingNotice,
			IsOpenWorkReply:    &defaultSettingNotice,
		}
		if _, err := uc.repo.AddSettingNotice(ctx, m); err != nil {
			return nil, errors2.WithStack(err)
		}
	}

	return m, nil
}

func (uc *CommunityUseCase) UpdateSettingNotice(ctx context.Context, req *v1.UpdateCommunitySettingNoticeRequest) error {
	m := &CommunitySettingNotice{
		Openid:             req.Openid,
		Unionid:            req.Unionid,
		IsOpenLikeWork:     &req.IsOpenLikeWork,
		IsOpenLikeComment:  &req.IsOpenLikeComment,
		IsOpenCommentReply: &req.IsOpenCommentReply,
		IsOpenWorkReply:    &req.IsOpenWorkReply,
	}
	m.ID = uint(req.Id)
	if err := uc.repo.UpdateSettingNotice(ctx, uint(req.Id), m); err != nil {
		return errors2.WithStack(err)
	}
	return nil
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

	article, _ := uc.repo.GetArticle(ctx, uint(req.ArticleId))
	if article != nil {
		go uc.sendTemplateMessage(req.ArticleId, CommunityNoticeTypeCommentWork, article.Openid, req.Openid, req.Content)

		if req.ReplyOpenid != "" && article.Openid != req.ReplyOpenid {
			go uc.sendTemplateMessage(req.ArticleId, CommunityNoticeTypeCommentReply, req.ReplyOpenid, req.Openid, req.Content)
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
	articles, err := uc.repo.GetArticleListByOpenid(ctx, req.Openid, 0, 1)
	if err != nil {
		return nil
	}
	if len(articles) > 0 && time.Now().Sub(articles[0].CreatedAt).Hours() < 1 {
		return errors.New(400, "", "发布频率高，请稍后再试")
	}

	for i, photo := range req.Photos {
		req.Photos[i].Url = uc.cosUC.TidyURLByPhoto(photo.Url)[0]
	}

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
			Openid:       openid,
			Username:     defaultCommentUser.Username,
			Avatar:       defaultCommentUser.Avatar,
			Introduction: defaultCommentUser.Introduction,
		}
		wechatUser, _ := uc.wechatUc.GetUser(ctx, appidCommunity, openid)
		if wechatUser != nil && len(wechatUser.Unionid) > 0 {
			item.Unionid = wechatUser.Unionid
		}
		_, err = uc.repo.AddUser(ctx, item)
		if err != nil {
			return nil, errors2.WithStack(err)
		}
	}
	item.Avatar, err = uc.cosUC.GetPresignedURL(ctx, item.Avatar)
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	return item, nil
}

func (uc *CommunityUseCase) UpdateUserUnionid(ctx context.Context, openid, unionid string) error {
	err := uc.repo.UpdateUserUnionid(ctx, openid, unionid)
	if err != nil {
		return errors2.WithStack(err)
	}
	return nil
}

func (uc *CommunityUseCase) UpdateMyProfile(ctx context.Context, req *v1.UpdateCommunityMyProfileRequest) error {
	m := &CommunityUser{
		Username:     req.Username,
		Introduction: req.Introduction,
		Avatar:       uc.cosUC.TidyURLByPhoto(req.AvatarUrl)[0],
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
	if page == 0 {
		topItems, errTop := uc.repo.GetTopArticleList(ctx)
		if errTop != nil {
			uc.log.Errorw("msg", "GetArticleList/GetTopArticleList", openid, "openid")
		} else {
			items = append(topItems, items...)
		}
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

func (uc *CommunityUseCase) UpdateLike(ctx context.Context, openid string, tid uint64, typ CommunityLikeType) (
	*v1.UpdateCommunityLikeReply, error,
) {
	like, err := uc.repo.GetLike(ctx, openid, uint(tid), typ)
	if err != nil && errors2.Cause(err) != gorm.ErrRecordNotFound {
		return nil, errors2.WithStack(err)
	}
	updateTo := CommunityLikeStatusLike
	isSendMsg := false
	if like == nil {
		isSendMsg = true
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
		if isSendMsg && step == 1 {
			article, _ := uc.repo.GetArticle(ctx, uint(tid))
			if article != nil {
				go uc.sendTemplateMessage(tid, CommunityNoticeTypeLikeWork, article.Openid, openid, "给你点个赞")
			}
		}
		err = uc.repo.UpdateArticleLikeCount(ctx, uint(tid), step)
		if err != nil {
			return nil, errors2.WithStack(err)
		}
	} else if typ == CommunityLikeTypeComment {
		if isSendMsg && step == 1 {
			comment, _ := uc.repo.GetComment(ctx, uint(tid))
			if comment != nil {
				go uc.sendTemplateMessage(comment.ArticleID, CommunityNoticeTypeLikeComment, comment.Openid, openid, "给你点个赞")
			}
		}
		err = uc.repo.UpdateCommentLikeCount(ctx, uint(tid), step)
		if err != nil {
			return nil, errors2.WithStack(err)
		}
	}

	return &v1.UpdateCommunityLikeReply{IsLike: updateTo == CommunityLikeStatusLike}, nil
}

func (uc *CommunityUseCase) sendTemplateMessage(articleID uint64, typ CommunityNoticeType, openid, sendOpenid, content string) {
	if openid == sendOpenid {
		return
	}
	ctx := context.Background()

	notice, err := uc.GetSettingNotice(ctx, openid)
	if err != nil {
		uc.log.Errorw("msg", "repo.GetSettingNotice", "openid", openid, "err", err)
		return
	}

	first := "你得到一个回复"
	switch typ {
	case CommunityNoticeTypeLikeComment:
		first = "矮油，你的评论不错哦"
		if !*notice.IsOpenLikeComment {
			return
		}
	case CommunityNoticeTypeLikeWork:
		first = "矮油，你的作品不错哦"
		if !*notice.IsOpenLikeWork {
			return
		}
	case CommunityNoticeTypeCommentWork:
		first = "你的作品有一个新的评论"
		if !*notice.IsOpenLikeComment {
			return
		}
	case CommunityNoticeTypeCommentReply:
		first = "你的评论被回复了"
		if !*notice.IsOpenCommentReply {
			return
		}
	}
	user, err := uc.wechatUc.GetUser(ctx, appidCommunity, openid)
	if err != nil {
		uc.log.Errorw("msg", "sendTemplateMessageGetUser", "openid", openid, "err", err)
		return
	}
	if user.Unionid == "" {
		return
	}
	sendUserUnionid := ""
	sendUser, _ := uc.wechatUc.GetUser(ctx, appidCommunity, sendOpenid)
	if sendUser != nil {
		sendUserUnionid = sendUser.Unionid
	}
	sendWxUser, err := uc.GetMyProfile(ctx, sendOpenid)
	if err != nil {
		uc.log.Errorw("msg", "sendTemplateMessageGetUser", "openid", openid, "sendOpenid", sendOpenid, "err", err)
		return
	}
	const templateID = "2YOC5fxi1h3KsxjCLq1Jea1zQzEabiwALahU4jPzkN4"
	var msg CommunityNoticeMessage

	data := make(map[string]*message.TemplateDataItem)
	data["first"] = &message.TemplateDataItem{Value: first}
	data["keyword1"] = &message.TemplateDataItem{Value: sendWxUser.Username}
	data["keyword2"] = &message.TemplateDataItem{Value: time.Now().Format("2006-01-02 15:04:05")}
	data["keyword3"] = &message.TemplateDataItem{Value: content}
	data["remark"] = &message.TemplateDataItem{Value: "点击卡片会跳转到文章。\n如需退订此类消息，请前往[你我他群]小程序，我的页面进行设置"}
	msg.TemplateMessage = message.TemplateMessage{
		TemplateID: templateID,
		Data:       data,
	}
	msg.MiniProgram.AppID = appidCommunity
	msg.MiniProgram.PagePath = fmt.Sprintf("pages/index/detail?id=%d", articleID)

	errMsg := "ok"
	msgID, err := uc.wechatUc.SendTemplateMsg(ctx, user.Unionid, &msg.TemplateMessage)
	if err != nil {
		errMsg = err.Error()
	}

	m := &CommunityNoticeHistory{
		Openid:     user.Openid,
		Unionid:    user.Unionid,
		ReqOpenid:  sendOpenid,
		ReqUnionid: sendUserUnionid,
		TemplateID: templateID,
		MsgID:      uint64(msgID),
		Detail:     msg,
		ErrMsg:     errMsg,
	}
	if _, err := uc.repo.AddNoticeHistory(ctx, m); err != nil {
		uc.log.Errorw("msg", "sendTemplateMessageSendTemplateMsg", "reqmsg", msg, "err", err)
		return
	}
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
		} else {
			avatar, err := uc.cosUC.GetPresignedURL(ctx, user.Avatar)
			if err != nil {
				return nil, errors2.WithStack(err)
			}
			user.Avatar = avatar
		}
		replyUser, found := userMap[item.ReplyOpenid]
		if !found {
			replyUser = defaultCommentUser
		} else {
			avatar, err := uc.cosUC.GetPresignedURL(ctx, replyUser.Avatar)
			if err != nil {
				return nil, errors2.WithStack(err)
			}
			replyUser.Avatar = avatar
		}
		_, isLike := likeMap[item.ID]

		if item.LikeCount < 0 {
			item.LikeCount = 0
		}
		if item.CommentCount < 0 {
			item.CommentCount = 0
		}

		userinfo := &v1.CommunityUser{
			Id:        uint64(user.ID),
			Openid:    user.Openid,
			Username:  user.Username,
			AvatarUrl: user.Avatar,
			TagValue:  user.TagValue,
			TagClass:  user.TagClass,
		}
		res[i] = &v1.AddCommunityCommentReply{
			Id:             uint64(item.ID),
			Content:        item.Content,
			Userinfo:       userinfo,
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
				TagValue:  replyUser.TagValue,
				TagClass:  replyUser.TagClass,
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
		} else {
			avatar, err := uc.cosUC.GetPresignedURL(ctx, user.Avatar)
			if err != nil {
				return nil, errors2.WithStack(err)
			}
			user.Avatar = avatar
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

		photos, err := uc.cosUC.GetPresignedURLByPhoto(ctx, item.Photos)
		if err != nil {
			return nil, errors2.WithStack(err)
		}
		res[i] = &v1.GetCommunityArticleReply{
			Id:              uint64(item.ID),
			PubUserName:     user.Username,
			PubUserAvatar:   user.Avatar,
			PubContent:      item.Content,
			PubUserTagValue: user.TagValue,
			PubUserTagClass: user.TagClass,
			PubTime:         item.CreatedAt.Format("01-02 15:04"),
			LikeCount:       uint64(item.LikeCount),
			ComCount:        uint64(item.CommentCount),
			IsLike:          isLike,
			Photos:          photos,
			Openid:          item.Openid,
			Type:            uint32(item.Type),
			MpAppid:         item.MpAppid,
			Url:             item.Url,
		}
	}

	return res, nil
}
