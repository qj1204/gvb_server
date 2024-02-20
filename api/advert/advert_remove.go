package advert

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
)

// AdvertRemoveView 删除广告
// @Tags 广告管理
// @Summary 删除广告
// @Description 删除广告
// @Param data body models.RemoveRequest true "查询参数"
// @Router /api/advert [delete]
// @Produce json
// @Success 200 {object} response.Response{data=string}
func (this *AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var advertList []models.AdvertModel
	count := global.DB.Find(&advertList, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("图片不存在", c)
		return
	}
	global.DB.Delete(&advertList)
	response.OkWithMessage(fmt.Sprintf("共删除%d条广告", count), c)
}
