package biz

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Appid     string `gorm:"type:char(18);index:idx_appid;"`
	Code      string `gorm:"type:varchar(20);index:idx_code;"`
	ImagePath string `gorm:"type:varchar(300);"`
	Title     string `gorm:"type:varchar(100);"`
	Content   string `gorm:"type:longtext;"`
	Sort      int    `gorm:"type:int(10);index:idx_sort;"`
}

func (Article) TableName() string {
	return "article"
}
