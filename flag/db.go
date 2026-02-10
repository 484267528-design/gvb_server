package flag

import (
	"gvb_server/global"
	"gvb_server/models"
)

func Makemigrations() {
	var err error
	global.DB.SetupJoinTable(&models.UserModel{}, "CollectModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.UserModel{},
			&models.UserCollectModel{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			&models.BannerModel{},
			&models.ArticleModel{},
			&models.AdvertModel{},
			&models.TagModel{},
			&models.CommentModel{},
			&models.FeedBackModel{},
			&models.MessageModel{},
			&models.LoginDataModel{},
		)
	if err != nil {
		global.Log.Error("[error] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功")
}
