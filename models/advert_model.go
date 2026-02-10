package models

// 广告
type AdvertModel struct {
	MODEL
	Title  string `gorm:"size:255" json:"title"` //标题
	Href   string `gorm:"size:255" json:"href"`  //跳转链接
	Image  string `gorm:"size:255" json:"image"`
	IsShow bool   ` json:"is_show"`
}
