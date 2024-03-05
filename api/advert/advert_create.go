package advert

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
)

type AdvertRequest struct {
	Title  string `gorm:"size:32" json:"title" binding:"required" msg:"请输入广告标题" structs:"title"` // 广告标题
	Href   string `json:"href" binding:"required,url" msg:"广告链接非法" structs:"href"`               // 广告链接
	Image  string `json:"image" binding:"required,url" msg:"广告图片地址" structs:"image"`             // 广告图片
	IsShow *bool  `json:"is_show" binding:"required" msg:"请选择是否展示" structs:"is_show"`            // 是否显示
}

// AdvertCreateView 添加广告
// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @Param data body AdvertRequest true "表示多个参数"
// @Router /api/advert [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (this *AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	// 重复广告判断
	var advert models.AdvertModel
	count := global.DB.Take(&advert, "title=?", cr.Title).RowsAffected
	if count > 0 {
		response.FailWithMessage("该广告已存在", c)
		return
	}

	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Image:  cr.Image,
		IsShow: *cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("添加广告失败", c)
		return
	}
	response.OkWithMessage("添加广告成功", c)
}
