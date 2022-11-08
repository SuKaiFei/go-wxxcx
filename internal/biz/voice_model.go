package biz

import (
	"database/sql/driver"
	"encoding/json"
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
	Appid         string     `gorm:"type:char(18);index:idx_appid;"`
	Code          string     `gorm:"type:varchar(50);index:idx_code;"`
	Name          string     `gorm:"type:varchar(50);"`
	Default       bool       `gorm:"type:tinyint;"`
	Type          VoiceType  `gorm:"type:int(1);default:1;"`
	MpAppid       string     `gorm:"type:char(18);"`
	Sort          int        `gorm:"type:int(1);default:1;"`
	ShareTitle    string     `gorm:"type:varchar(50);"`
	ShareImageUrl string     `gorm:"type:varchar(300);"`
	Works         VoiceWorks `gorm:"column:works;type:json;"`
}

func (Voice) TableName() string {
	return "voice"
}

func (m VoiceWorks) Value() (driver.Value, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *VoiceWorks) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), m)
}