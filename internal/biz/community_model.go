package biz

import (
	"database/sql/driver"
	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	jsoniter "github.com/json-iterator/go"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"gorm.io/gorm"
)

type CommunityNoticeType uint8

const (
	CommunityNoticeTypeLikeWork CommunityNoticeType = iota + 1
	CommunityNoticeTypeLikeComment
	CommunityNoticeTypeCommentReply
	CommunityNoticeTypeCommentWork
)

type CommunitySettingNotice struct {
	gorm.Model
	Openid             string `gorm:"type:char(28);"`
	Unionid            string `gorm:"type:char(28);"`
	IsOpenLikeWork     *bool
	IsOpenLikeComment  *bool
	IsOpenCommentReply *bool
	IsOpenWorkReply    *bool
}

func (CommunitySettingNotice) TableName() string {
	return "community_setting_notice"
}

type CommunityNoticeMessage struct {
	message.TemplateMessage `json:",inline"`
}

func (m CommunityNoticeMessage) Value() (driver.Value, error) {
	return jsoniter.MarshalToString(m)
}

func (m *CommunityNoticeMessage) Scan(input interface{}) error {
	return jsoniter.Unmarshal(input.([]byte), m)
}

type CommunityNoticeHistory struct {
	gorm.Model
	Openid     string                 `gorm:"type:char(28);"`
	Unionid    string                 `gorm:"type:char(28);"`
	ReqOpenid  string                 `gorm:"type:char(28);"`
	ReqUnionid string                 `gorm:"type:char(28);"`
	TemplateID string                 `gorm:"type:char(43);"`
	MsgID      uint64                 `gorm:"type:bigint"`
	Detail     CommunityNoticeMessage `gorm:"type:json"`
	ErrMsg     string                 `gorm:"type:varchar(100);"`
}

func (CommunityNoticeHistory) TableName() string {
	return "community_notice_history"
}

type CommunityUser struct {
	gorm.Model
	Openid   string `gorm:"type:char(28);"`
	Unionid  string `gorm:"type:char(28);"`
	Username string `gorm:"type:varchar(50);"`
	Avatar   string `gorm:"type:varchar(500);"`
}

func (CommunityUser) TableName() string {
	return "community_user"
}

type CommunityArticlePhotos []*v1.Photo

func (m CommunityArticlePhotos) Value() (driver.Value, error) {
	return jsoniter.MarshalToString(m)
}

func (m *CommunityArticlePhotos) Scan(input interface{}) error {
	return jsoniter.Unmarshal(input.([]byte), m)
}

type CommunityArticle struct {
	gorm.Model
	Openid       string                 `gorm:"type:char(28);"`
	Content      string                 `gorm:"type:longtext;"`
	IsTop        bool                   `gorm:"default:false"`
	Photos       CommunityArticlePhotos `gorm:"type:json;"`
	CommentCount int64                  `gorm:"type:bigint;default:0;"`
	LikeCount    int64                  `gorm:"type:bigint;default:0;"`
	Type         NavigationType         `gorm:"type:int(1);default:1;"`
	MpAppid      string                 `gorm:"type:char(18);"`
	Url          string                 `gorm:"type:varchar(300);"`
}

func (CommunityArticle) TableName() string {
	return "community_article"
}

type CommunityComment struct {
	gorm.Model
	ArticleID        uint64 `gorm:"type:bigint;"`
	Openid           string `gorm:"type:char(28);"`
	ReplyOpenid      string `gorm:"type:char(28);"`
	TopReplyOpenid   string `gorm:"type:char(28);"`
	Content          string `gorm:"type:longtext;"`
	LikeCount        int64  `gorm:"type:bigint;default:0;"`
	CommentCount     int64  `gorm:"type:bigint;default:0;"`
	InteractiveCount uint64 `gorm:"type:bigint;default:0;"`
	ReplyID          uint64 `gorm:"type:bigint;"`
	TopReplyID       uint64 `gorm:"type:bigint;default:0;"`
}

func (CommunityComment) TableName() string {
	return "community_comment"
}

type CommunityLikeType uint8
type CommunityLikeStatus uint8

const (
	CommunityLikeTypeComment CommunityLikeType = iota + 1
	CommunityLikeTypeArticle
)

const (
	CommunityLikeStatusLike CommunityLikeStatus = iota + 1
	CommunityLikeStatusUnlike
)

type CommunityLike struct {
	gorm.Model
	Openid string              `gorm:"type:char(28);"`
	Type   CommunityLikeType   `gorm:"type:tinyint;"`
	Tid    uint64              `gorm:"type:bigint;"`
	Status CommunityLikeStatus `gorm:"type:tinyint;"`
}

func (CommunityLike) TableName() string {
	return "community_like"
}

type CommunityFeedback struct {
	gorm.Model
	Openid  string            `gorm:"type:char(28);"`
	Type    CommunityLikeType `gorm:"type:tinyint;"`
	Tid     uint64            `gorm:"type:bigint;"`
	Content string            `gorm:"type:longtext;"`
}

func (CommunityFeedback) TableName() string {
	return "community_feedback"
}
