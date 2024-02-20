package menu

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

func (this *MenuApi) MenuListView(c *gin.Context) {
	// 先查菜单
	var menuList []models.MenuModel
	var menuIDList []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)

	// 查menu_banner连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in (?)", menuIDList)
	// menuBanners
	//{1 24 {{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}    [] 0 [] 0 0} {{24 2023-12-20 16:48:57.199 +0800 CST 2023-12-20 16:48:57.199 +0800 CST} static/uploads/伊蕾娜1.jpeg a52aec83ff54878baec0c65dab7b1ac6 伊蕾娜1.jpeg 本地} 2}
	//{1 23 {{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}    [] 0 [] 0 0} {{23 2023-12-20 16:48:57.18 +0800 CST 2023-12-20 16:48:57.18 +0800 CST} static/uploads/武汉.jpg e67799d719a50eebd4c94d9981806441 武汉.jpg 本地} 1}
	//{1 22 {{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}    [] 0 [] 0 0} {{22 2023-12-20 16:47:47.336 +0800 CST 2023-12-20 16:47:47.336 +0800 CST} http://s5x2fcgbb.hn-bkt.clouddn.com/gvb/20231220164747_google.png 14b07f21274b69891b1ad6c8a6a1fb05 google.png 七牛云} 0}
	var menuResponseList []MenuResponse
	for _, menu := range menuList {
		// 解决nil值的问题
		var banners = make([]Banner, 0)
		for _, menuBanner := range menuBanners {
			if menu.ID != menuBanner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   menuBanner.BannerID,
				Path: menuBanner.BannerModel.Path,
			})
		}
		menuResponseList = append(menuResponseList, MenuResponse{
			MenuModel: menu,
			Banners:   banners,
		})
	}
	response.OkWithData(menuResponseList, c)
}
