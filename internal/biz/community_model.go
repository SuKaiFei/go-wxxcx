package biz

import (
	"database/sql/driver"
	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
)

type CommunityUser struct {
	gorm.Model
	Openid   string `gorm:"type:char(28);"`
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
	Photos       CommunityArticlePhotos `gorm:"type:json;"`
	CommentCount int64                  `gorm:"type:bigint;default:0;"`
	LikeCount    int64                  `gorm:"type:bigint;default:0;"`
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
