package data

import (
	"context"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	errors2 "github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
)

type CommunityRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommunityRepo(data *Data, logger log.Logger) biz.CommunityRepo {
	return &CommunityRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *CommunityRepo) GetSendUserTitle(ctx context.Context) ([]string, error) {
	var tmp []struct {
		ArticleId int    `gorm:"column:article_id;"`
		Openid    string `gorm:"column:aopenid;"`
		Ac        int    `gorm:"column:ac;"`
	}
	err := r.data.db.WithContext(ctx).Raw(`SELECT
	article_id,
	aopenid,
	count( article_id ) ac 
FROM
	(
	SELECT
		article_id,
		cc.openid copenid,
		ca.openid aopenid 
	FROM
		community_comment cc
		LEFT JOIN community_article ca ON cc.article_id = ca.id 
	WHERE
		ca.created_at > DATE_SUB( curdate(), INTERVAL 1 DAY ) 
	GROUP BY
		article_id,
		cc.openid,
		ca.openid 
	) s 
GROUP BY
	article_id,
	aopenid 
ORDER BY
	ac DESC LIMIT 10
	`).Scan(&tmp).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	res := make([]string, len(tmp))
	for i, s := range tmp {
		res[i] = s.Openid
	}
	return res, nil
}

func (r *CommunityRepo) GetUserTitleByValue(ctx context.Context, openid, value string) (m *biz.CommunityUserTitle, err error) {
	err = r.data.db.WithContext(ctx).First(&m, "openid=? and value=?", openid, value).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return
}

func (r *CommunityRepo) AddUserTitle(ctx context.Context, m *biz.CommunityUserTitle) (uint, error) {
	err := r.data.db.WithContext(ctx).Create(m).Error
	if err != nil {
		return 0, errors2.WithStack(err)
	}
	return m.ID, nil
}

func (r *CommunityRepo) GetUserTitle(ctx context.Context, id uint) (m *biz.CommunityUserTitle, err error) {
	err = r.data.db.WithContext(ctx).First(&m, "id=?", id).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return m, nil
}

func (r *CommunityRepo) UpdateUserTitle(ctx context.Context, id uint, m *biz.CommunityUserTitle) error {
	err := r.data.db.WithContext(ctx).Where("id=?", id).Updates(&m).Error
	if err != nil {
		return errors2.WithStack(err)
	}
	return nil
}

func (r *CommunityRepo) GetUserTitleList(ctx context.Context, openid string) (m []*biz.CommunityUserTitle, err error) {
	err = r.data.db.WithContext(ctx).Find(&m, "openid=? and validity_period>now()", openid).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return m, nil
}

func (r *CommunityRepo) GetSettingNotice(ctx context.Context, openid string) (m *biz.CommunitySettingNotice, err error) {
	err = r.data.db.WithContext(ctx).First(&m, "openid=?", openid).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return m, nil
}

func (r *CommunityRepo) AddSettingNotice(ctx context.Context, m *biz.CommunitySettingNotice) (uint, error) {
	err := r.data.db.WithContext(ctx).Create(m).Error
	if err != nil {
		return 0, errors2.WithStack(err)
	}
	return m.ID, nil
}

func (r *CommunityRepo) AddNoticeHistory(ctx context.Context, m *biz.CommunityNoticeHistory) (uint, error) {
	err := r.data.db.WithContext(ctx).Create(m).Error
	if err != nil {
		return 0, errors2.WithStack(err)
	}
	return m.ID, nil
}

func (r *CommunityRepo) UpdateSettingNotice(ctx context.Context, id uint, m *biz.CommunitySettingNotice) error {
	err := r.data.db.WithContext(ctx).Where("id=?", id).Updates(m).Error
	if err != nil {
		return errors2.WithStack(err)
	}
	return nil
}

func (r *CommunityRepo) GetLike(ctx context.Context, openid string, tid uint, likeType biz.CommunityLikeType) (
	m *biz.CommunityLike, err error,
) {
	err = r.data.db.WithContext(ctx).
		Where("openid=? and tid=? and type=?", openid, tid, likeType).
		First(&m).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return
}

func (r *CommunityRepo) AddLike(ctx context.Context, m *biz.CommunityLike) (uint, error) {
	err := r.data.db.WithContext(ctx).
		Create(&m).Error
	if err != nil {
		return 0, errors2.WithStack(err)
	}
	return m.ID, nil
}

func (r *CommunityRepo) AddArticle(ctx context.Context, m *biz.CommunityArticle) (uint, error) {
	err := r.data.db.WithContext(ctx).
		Create(&m).Error
	if err != nil {
		return 0, errors2.WithStack(err)
	}
	return m.ID, nil
}

func (r *CommunityRepo) AddComment(ctx context.Context, m *biz.CommunityComment) (uint, error) {
	err := r.data.db.WithContext(ctx).
		Create(&m).Error
	if err != nil {
		return 0, errors2.WithStack(err)
	}
	return m.ID, nil
}

func (r *CommunityRepo) UpdateLikeStatus(ctx context.Context, id uint, status biz.CommunityLikeStatus) error {
	err := r.data.db.WithContext(ctx).Model(new(biz.CommunityLike)).Limit(1).
		Where("id=?", id).
		Update("status", status).Error
	if err != nil {
		return errors2.WithStack(err)
	}
	return nil
}

func (r *CommunityRepo) GetUserListByOpenid(ctx context.Context, openIDs []string) (m []*biz.CommunityUser, err error) {
	err = r.data.db.WithContext(ctx).Find(&m, "openid in ?", openIDs).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return
}

func (r *CommunityRepo) GetArticleList(ctx context.Context, page, pageSize int) (m []*biz.CommunityArticle, err error) {
	err = r.data.db.WithContext(ctx).Order("id desc").Limit(pageSize).Offset(page * pageSize).Find(&m).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return
}

func (r *CommunityRepo) GetTopArticleList(ctx context.Context) (m []*biz.CommunityArticle, err error) {
	err = r.data.db.WithContext(ctx).Where("is_top=?", true).Find(&m).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return
}

func (r *CommunityRepo) GetCommentList(ctx context.Context, aid, comID uint, page, pageSize int) (
	m []*biz.CommunityComment, err error,
) {
	db := r.data.db.WithContext(ctx).Order("id desc")
	db = db.Where("article_id = ?", aid)
	db = db.Where("top_reply_id = ?", comID)
	err = db.Limit(pageSize).Offset(page * pageSize).Find(&m).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return
}

func (r *CommunityRepo) GetComment(ctx context.Context, id uint) (m *biz.CommunityComment, err error) {
	err = r.data.db.WithContext(ctx).First(&m, "id=?", id).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return
}

func (r *CommunityRepo) GetArticleListByOpenid(ctx context.Context, openid string, page, pageSize int) (m []*biz.CommunityArticle, err error) {
	err = r.data.db.WithContext(ctx).
		Where("openid=?", openid).Order("id desc").Limit(pageSize).Offset(page * pageSize).
		Find(&m).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return
}

func (r *CommunityRepo) GetArticle(ctx context.Context, id uint) (m *biz.CommunityArticle, err error) {
	err = r.data.db.WithContext(ctx).First(&m, "id=?", id).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return
}

func (r *CommunityRepo) GetUser(ctx context.Context, openid string) (m *biz.CommunityUser, err error) {
	err = r.data.db.WithContext(ctx).First(&m, "openid=?", openid).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return
}

func (r *CommunityRepo) AddUser(ctx context.Context, m *biz.CommunityUser) (uint, error) {
	err := r.data.db.WithContext(ctx).Create(m).Error
	if err != nil {
		return 0, errors2.WithStack(err)
	}
	return m.ID, nil
}

func (r *CommunityRepo) AddFeedback(ctx context.Context, m *biz.CommunityFeedback) (uint, error) {
	err := r.data.db.WithContext(ctx).Create(m).Error
	if err != nil {
		return 0, errors2.WithStack(err)
	}
	return m.ID, nil
}

func (r *CommunityRepo) DeleteArticle(ctx context.Context, openid string, id uint) (int64, error) {
	db := r.data.db.WithContext(ctx).Model(new(biz.CommunityArticle)).
		Delete("id=? and openid=?", id, openid)
	err := db.Error
	if err != nil {
		return 0, errors2.WithStack(err)
	}
	return db.RowsAffected, nil
}

func (r *CommunityRepo) DeleteComment(ctx context.Context, openid string, id uint) (int64, error) {
	db := r.data.db.WithContext(ctx).Model(new(biz.CommunityComment)).
		Delete("id=? and openid=?", id, openid)
	err := db.Error
	if err != nil {
		return 0, errors2.WithStack(err)
	}
	return db.RowsAffected, nil
}

func (r *CommunityRepo) UpdateUser(ctx context.Context, id uint, m *biz.CommunityUser) error {
	values := map[string]interface{}{
		"username":            m.Username,
		"avatar":              m.Avatar,
		"introduction":        m.Introduction,
		"tag_id":              m.TagID,
		"tag_value":           m.TagValue,
		"tag_class":           m.TagClass,
		"tag_validity_period": m.TagValidityPeriod,
	}
	err := r.data.db.WithContext(ctx).Where("id = ?", id).Model(m).Updates(values).Error
	if err != nil {
		return errors2.WithStack(err)
	}
	return nil
}

func (r *CommunityRepo) UpdateUserUserTitle(ctx context.Context, m *biz.CommunityUserTitle) error {
	values := map[string]interface{}{
		"tag_id":              m.ID,
		"tag_value":           m.Value,
		"tag_class":           m.Class,
		"tag_validity_period": m.ValidityPeriod,
	}
	err := r.data.db.WithContext(ctx).Where("openid=?", m.Openid).Model(new(biz.CommunityUser)).Updates(values).Error
	if err != nil {
		return errors2.WithStack(err)
	}
	return nil
}

func (r *CommunityRepo) UpdateUserUnionid(ctx context.Context, openid, unionid string) error {
	err := r.data.db.WithContext(ctx).Model(new(biz.CommunityUser)).
		Where("openid = ?", openid).Update("unionid", unionid).Error
	if err != nil {
		return errors2.WithStack(err)
	}
	return nil
}

func (r *CommunityRepo) GetArticleLikeByOpenid(ctx context.Context, openid string, articles []uint) (
	m []*biz.CommunityLike, err error,
) {
	err = r.data.db.WithContext(ctx).
		Where("type=? and openid=? and tid in ?", biz.CommunityLikeTypeArticle, openid, articles).
		Limit(len(articles)).Find(&m).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return
}

func (r *CommunityRepo) GetCommentLikeByOpenid(ctx context.Context, openid string, articles []uint) (
	m []*biz.CommunityLike, err error,
) {
	err = r.data.db.WithContext(ctx).
		Where("type=? and openid=? and tid in ?", biz.CommunityLikeTypeComment, openid, articles).
		Limit(len(articles)).Find(&m).Error
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return
}

func (r *CommunityRepo) UpdateArticleLikeCount(ctx context.Context, id uint, step int) error {
	err := r.data.db.WithContext(ctx).Model(new(biz.CommunityArticle)).
		Where("id=?", id).
		Update("like_count", gorm.Expr("like_count + ?", step)).Error
	if err != nil {
		return errors2.WithStack(err)
	}
	return nil
}

func (r *CommunityRepo) UpdateArticleCommentCount(ctx context.Context, id uint, step int) error {
	err := r.data.db.WithContext(ctx).Model(new(biz.CommunityArticle)).
		Where("id=?", id).
		Update("comment_count", gorm.Expr("comment_count + ?", step)).Error
	if err != nil {
		return errors2.WithStack(err)
	}
	return nil
}

func (r *CommunityRepo) UpdateCommentLikeCount(ctx context.Context, id uint, step int) error {
	err := r.data.db.WithContext(ctx).Model(new(biz.CommunityComment)).Debug().
		Where("id=?", id).
		Updates(map[string]interface{}{
			"like_count":        gorm.Expr("like_count + ?", step),
			"interactive_count": gorm.Expr("interactive_count + ?", step),
		}).Error
	if err != nil {
		return errors2.WithStack(err)
	}
	return nil
}

func (r *CommunityRepo) UpdateCommentCount(ctx context.Context, id uint, step int) error {
	err := r.data.db.WithContext(ctx).Model(new(biz.CommunityComment)).
		Where("id=?", id).
		Updates(map[string]interface{}{
			"comment_count":     gorm.Expr("comment_count + ?", step),
			"interactive_count": gorm.Expr("interactive_count + ?", step),
		}).Error
	if err != nil {
		return errors2.WithStack(err)
	}
	return nil
}
