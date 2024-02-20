package advert

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
)

// AdvertUpdateView 更新广告
// @Tags 广告管理
// @Summary 更新广告
// @Description 更新广告
// @Param data body AdvertRequest true "广告的一些参数"
// @Router /api/advert/:id [put]
// @Produce json
// @Success 200 {object} response.Response{data=string}
func (*AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	// 判断广告是否存在
	var advert models.AdvertModel
	count := global.DB.Take(&advert, id).RowsAffected
	if count == 0 {
		response.FailWithMessage("广告不存在", c)
		return
	}

	// 结构体转map的第三方包structs
	maps := structs.Map(&cr)
	err = global.DB.Model(&advert).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("修改广告失败", c)
		return
	}
	response.OkWithMessage("修改广告成功", c)
}
