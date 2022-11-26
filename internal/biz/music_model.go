package biz

import "gorm.io/gorm"

type Music struct {
	gorm.Model
	Appid         string `gorm:"type:char(18);"`
	Code          string `gorm:"type:varchar(20);"`
	ImagePath     string `gorm:"type:varchar(300);"`
	Name          string `gorm:"type:varchar(100);"`
	Singer        string `gorm:"type:varchar(100);"`
	Url           string `gorm:"type:varchar(200);"`
	Duration      uint64 `gorm:"type:int(10);"`
	ShareTitle    string `gorm:"type:varchar(50);"`
	ShareImageUrl string `gorm:"type:varchar(300);"`
	Sort          int    `gorm:"type:int(10);"`
}

func (Music) TableName() string {
	return "music"
}
