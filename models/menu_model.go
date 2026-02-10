package models

import "gvb_server/models/ctype"

type MenuModel struct {
	MODEL
	MenuTitle    string        `gorm:"size:32" json:"menu_title"`
	MenuTitleEn  string        `gorm:"size:32" json:"menu_title_en"json:"menu_title_en"`
	Slogan       string        `gorm:"size:32" json:"slogan"`
	Abstract     ctype.Array   `gorm:"type:string;size:64" json:"abstract"`                                                       //简介
	AbstractTime int64         ` json:"abstract_time"`                                                                            //简介切换时间
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;joinReferences:BannerID" json:"banners"` //菜单图片列表
	BannerTime   int           ` json:"banner_time"`                                                                              //菜单图片切换时间
	Sort         int           `gorm:"size:10" json:"sort"`                                                                       //菜单顺序
}
