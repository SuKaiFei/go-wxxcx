package biz

import "gorm.io/gorm"

type NavigationType uint8

const (
	NavigationTypeInner NavigationType = iota + 1 // 内部页面
	NavigationTypeMp                              // 打开小程序
	NavigationTypeImage                           // 打开图片
)

type Navigation struct {
	gorm.Model
	Appid     string         `gorm:"type:char(18);"`
	Code      string         `gorm:"type:varchar(20);"`
	Type      NavigationType `gorm:"type:int(1);default:1;"`
	ImagePath string         `gorm:"type:varchar(300);"`
	Title     string         `gorm:"type:varchar(100);"`
	Describe  string         `gorm:"type:varchar(300);"`
	Sort      int            `gorm:"type:int(10);"`
	MpAppid   string         `gorm:"type:char(18);"`
	Url       string         `gorm:"type:varchar(300);"`
}

func (Navigation) TableName() string {
	return "navigation"
}
