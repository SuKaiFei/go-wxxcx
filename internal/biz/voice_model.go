package biz

import (
	"database/sql/driver"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
)

type VoiceType uint8

const (
	VoiceTypeInner VoiceType = iota + 1 // 内部语音
	VoiceTypeMp                         // 打开小程序
)

type VoiceWork struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	TpUrl string `json:"tp_url"`
}

type VoiceWorks [][]*VoiceWork

type Voice struct {
	gorm.Model
	Appid         string     `gorm:"type:char(18);"`
	Code          string     `gorm:"type:varchar(50);"`
	Name          string     `gorm:"type:varchar(50);"`
	Default       bool       `gorm:"type:tinyint;"`
	Type          VoiceType  `gorm:"type:int(1);default:1;"`
	MpAppid       string     `gorm:"type:char(18);"`
	MpUrl         string     `gorm:"type:varchar(300);;"`
	Sort          int        `gorm:"type:int(1);default:1;"`
	ShareTitle    string     `gorm:"type:varchar(50);"`
	ShareImageUrl string     `gorm:"type:varchar(300);"`
	Works         VoiceWorks `gorm:"type:json;"`
}

func (Voice) TableName() string {
	return "voice"
}

func (m VoiceWorks) Value() (driver.Value, error) {
	return jsoniter.MarshalToString(m)
}

func (m *VoiceWorks) Scan(input interface{}) error {
	return jsoniter.Unmarshal(input.([]byte), m)
}
