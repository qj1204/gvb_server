package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
)

func init() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
}
func main() {
	var menu models.MenuModel
	global.DB.Take(&menu, 1)

	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	global.DB.Model(&menu).Association("Banners").Clear()
	var banners []models.MenuBannerModel

	type MenuImageID struct {
		ImageID uint `json:"image_id"`
		Sort    int  `json:"sort"`
	}

	var menus = []MenuImageID{
		{ImageID: 14, Sort: 1},
		{ImageID: 15, Sort: 2},
	}
	for _, img := range menus {
		banners = append(banners, models.MenuBannerModel{
			MenuID:   menu.ID,
			BannerID: img.ImageID,
			Sort:     img.Sort,
		})
	}
	global.DB.Create(&banners)

}
