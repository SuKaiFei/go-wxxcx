package biz

import (
	"gorm.io/gorm"
)

type BiaoQingBaoIndex struct {
	gorm.Model
	Appid     string `json:"app_id" gorm:"type:char(18);index:idx_appid;"`
	Name      string `json:"name" gorm:"type:varchar(255);"`
	Type      string `json:"type" gorm:"type:varchar(255);index:idx_type;"`
	Sort      int    `json:"sort" gorm:"type:int(255);index:idx_sort;"`
	ImagePath string `json:"image_path" gorm:"type:varchar(255);"`
}

func (BiaoQingBaoIndex) TableName() string {
	return "bqb_index"
}

type BiaoQingBao struct {
	gorm.Model
	Appid     string `json:"app_id" gorm:"type:char(18);index:idx_appid;"`
	Type      string `json:"type" gorm:"type:varchar(255);index:idx_type;"`
	ImagePath string `json:"image_path" gorm:"type:varchar(255);"`
}

func (BiaoQingBao) TableName() string {
	return "bqb"
}

type BiaoQingBaoIndexNum struct {
	Type string
	Num  uint64
}
