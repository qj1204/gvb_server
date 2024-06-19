package menu

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/response"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`                // 切换的时间，单位：秒
	BannerTime    int         `json:"banner_time" structs:"banner_time"`                    // 切换的时间，单位：秒
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号" structs:"sort"` // 菜单的顺序
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`                          // 具体图片的顺序
}

// MenuCreateView 发布菜单
// @Tags 菜单管理
// @Summary 发布菜单
// @Description 发布菜单
// @Param data body MenuRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/menus [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	// 重复值判断
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, "title = ? or path = ?", cr.Title, cr.Path).RowsAffected
	if count > 0 {
		response.FailWithMessage("菜单名称或者路径重复", c)
		return
	}

	// 创建数据，写入menu表
	menuModel := &models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(menuModel).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("菜单添加失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		response.OkWithMessage("菜单添加成功", c)
		return
	}

	var menuBannerList []models.MenuBannerModel
	for _, image := range cr.ImageSortList {
		// 这里需要判断imageSort.ImageID是否存在
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: image.ImageID,
			Sort:     image.Sort,
		})
	}
	// 写入menu_banner连接表
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("菜单添加失败", c)
		return
	}
	response.OkWithMessage("菜单添加成功", c)
}
