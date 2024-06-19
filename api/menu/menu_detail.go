package menu

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
)

// MenuDetailView 菜单详情
// @Tags 菜单管理
// @Summary 菜单详情
// @Description 菜单详情
// @Param id path int  true  "id"
// @Router /api/menus/{id} [get]
// @Produce json
// @Success 200 {object} response.Response{data=MenuResponse}
func (MenuApi) MenuDetailView(c *gin.Context) {
	id := c.Param("id")
	// 先查菜单
	var menuModel models.MenuModel
	err := global.DB.Take(&menuModel, id).Error
	if err != nil {
		response.FailWithMessage("菜单不存在", c)
		return
	}

	// 查menu_banner连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", id)
	// menuBanners
	//{1 24 {{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}    [] 0 [] 0 0} {{24 2023-12-20 16:48:57.199 +0800 CST 2023-12-20 16:48:57.199 +0800 CST} static/uploads/伊蕾娜1.jpeg a52aec83ff54878baec0c65dab7b1ac6 伊蕾娜1.jpeg 本地} 2}
	//{1 23 {{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}    [] 0 [] 0 0} {{23 2023-12-20 16:48:57.18 +0800 CST 2023-12-20 16:48:57.18 +0800 CST} static/uploads/武汉.jpg e67799d719a50eebd4c94d9981806441 武汉.jpg 本地} 1}
	//{1 22 {{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}    [] 0 [] 0 0} {{22 2023-12-20 16:47:47.336 +0800 CST 2023-12-20 16:47:47.336 +0800 CST} http://s5x2fcgbb.hn-bkt.clouddn.com/gvb/20231220164747_google.png 14b07f21274b69891b1ad6c8a6a1fb05 google.png 七牛云} 0}

	// 解决nil值的问题
	var banners = make([]Banner, 0)
	for _, menuBanner := range menuBanners {
		if menuModel.ID != menuBanner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   menuBanner.BannerID,
			Path: menuBanner.BannerModel.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}

	response.OkWithData(menuResponse, c)

}
