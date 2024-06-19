package menu

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
)

// MenuUpdateView 更新菜单
// @Tags 菜单管理
// @Summary 更新菜单
// @Description 更新菜单
// @Param data body MenuRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Param id path int  true  "id"
// @Router /api/menus/{id} [put]
// @Produce json
// @Success 200 {object} response.Response{}
func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")
	// 判断菜单是否存在
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		response.FailWithMessage("菜单不存在", c)
		return
	}

	// 先把之间的menu_banner清空
	global.DB.Model(&menuModel).Association("Banners").Clear()

	// 如果选择了banner，那就添加banner
	if len(cr.ImageSortList) > 0 {
		var menuBannerList []models.MenuBannerModel
		for _, image := range cr.ImageSortList {
			menuBannerList = append(menuBannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: image.ImageID,
				Sort:     image.Sort,
			})
		}
		err = global.DB.Create(&menuBannerList).Error
		if err != nil {
			global.Log.Error(err)
			response.FailWithMessage("菜单更新图片失败", c)
			return
		}
	}

	// 普通更新
	maps := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("菜单更新失败", c)
		return
	}
	response.OkWithMessage("菜单更新成功", c)
}
