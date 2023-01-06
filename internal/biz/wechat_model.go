package biz

import (
	"gorm.io/gorm"
)

type WechatUser struct {
	gorm.Model
	Appid   string `gorm:"type:char(18);"`
	Openid  string `gorm:"type:char(28);"`
	Unionid string `gorm:"type:char(28);"`
}

func (WechatUser) TableName() string {
	return "wechat_user"
}

type WechatOAUser struct {
	gorm.Model
	Openid  string `gorm:"type:char(28);"`
	Unionid string `gorm:"type:char(28);"`

	Subscribe      int32
	Nickname       string `gorm:"type:varchar(100);"`
	Sex            int32
	City           string `gorm:"type:varchar(100);"`
	Country        string `gorm:"type:varchar(100);"`
	Province       string `gorm:"type:varchar(100);"`
	Language       string `gorm:"type:varchar(50);"`
	Headimgurl     string `gorm:"type:varchar(200);"`
	SubscribeTime  int32  `gorm:"type:int;"`
	Remark         string `gorm:"type:varchar(500);"`
	GroupID        int32
	TagIDList      string `gorm:"type:varchar(500);"`
	SubscribeScene string `gorm:"type:varchar(50);"`
	QrScene        int
	QrSceneStr     string `gorm:"type:varchar(100);"`
}

func (WechatOAUser) TableName() string {
	return "wechat_oauser"
}
