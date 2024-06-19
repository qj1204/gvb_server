package flags

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/plugins/log_stash"
	log_stash_v2 "gvb_server/plugins/log_stash_v2"
)

func DB() {
	var err error
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	// 生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},
			&models.TagModel{},
			&models.MessageModel{},
			&models.AdvertModel{},
			&models.UserModel{},
			&models.CommentModel{},
			&models.UserCollectModel{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			&models.LoginDataModel{},
			&models.ChatModel{},
			&models.FeedbackModel{},
			&log_stash.LogStashModel{},
			&log_stash_v2.LogModel{},
			&models.UserScopeModel{},
			&models.AutoReplyModel{},
			&models.BigModelTagModel{},     // 大模型标签表
			&models.BigModelRoleModel{},    // 大模型角色表
			&models.BigModelChatModel{},    // 大模型对话表
			&models.BigModelSessionModel{}, // 大模型会话表
		)
	if err != nil {
		global.Log.Error("[error] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[success] 生成数据库表结构成功！")
}
