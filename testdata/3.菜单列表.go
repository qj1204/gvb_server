package main

import (
	"encoding/json"
	"fmt"
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

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type Menus struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

func main() {
	var list []models.MenuModel
	var menuIDList []uint
	// 获取菜单的 id 列表
	global.DB.Order("sort desc").Find(&list).Select("id").Scan(&menuIDList)
	// 根据menu_id查连接表
	var menuImages []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort asc").Find(&menuImages, "menu_id in ?", menuIDList)
	var menuList []Menus
	for _, model := range list {
		// 找菜单的图片列表
		var images []Banner
		for _, image := range menuImages {
			// 有的菜单没有图片
			if model.ID != image.MenuID {
				continue
			}
			images = append(images, Banner{
				ID:   image.BannerID,
				Path: image.BannerModel.Path,
			})
		}
		menuList = append(menuList, Menus{
			MenuModel: model,
			Banners:   images,
		})
	}
	byteData, err := json.Marshal(menuList)
	fmt.Println(string(byteData), err)
}
