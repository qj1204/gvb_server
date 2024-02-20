package flag

import (
	"gvb_server/global"
	"gvb_server/models"
)

func MakeMigrations() {
	var err error
	//global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})

	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.UserModel{},
		&models.MenuModel{},
		&models.BannerModel{},
		&models.TagModel{},
		//&models.ArticleModel{},
		&models.MessageModel{},
		&models.AdvertModel{},
		&models.CommentModel{},
		&models.MenuBannerModel{},
		&models.FeedbackModel{},
		&models.LoginDataModel{},
	)
	if err != nil {
		global.Log.Error("[error] 生成数据库表失败")
		return
	}
	global.Log.Info("[success] 生成数据库表成功")
}
