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

	// 根据menuIDList的id查menu_banner连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in (?)", menuIDList)
	// menuBanners
	/*[
	{1 5 {{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}    [] 0 [] 0 0} {{5 2024-02-27 16:52:44.323 +0800 CST 2024-02-27 16:52:44.323 +0800 CST} http://qiniu.xiaoxinqj.top/gvb/20240227165244_a5.jpg 15365249953e302b36f539e29f37f9aa a5.jpg 七牛云} 1}
	{1 4 {{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}    [] 0 [] 0 0} {{4 2024-02-27 16:52:44.239 +0800 CST 2024-02-27 16:52:44.239 +0800 CST} http://qiniu.xiaoxinqj.top/gvb/20240227165243_a4.jpg 7822a6008d890a2357aff37955478d3b a4.jpg 七牛云} 0}
	]*/
	var menuResponseList []MenuResponse
	for _, menu := range menuList {
		// 解决null值的问题
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

	// 用原生sql查询

	response.OkWithData(menuResponseList, c)
}
