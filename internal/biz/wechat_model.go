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
