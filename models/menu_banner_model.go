package models

// 自定义菜单和背景图的连接表
type MenuBannerModel struct {
	MenuID      int         `json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID"`
	BannerID    int         `json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID"`
	Sort        int         `gorm:"size:10" json:"sort"`
}
