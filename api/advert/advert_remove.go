package advert

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
)

// AdvertRemoveView 批量删除广告
// @Tags 广告管理
// @Summary 批量删除广告
// @Description 批量删除广告
// @Param data body models.RemoveRequest true "广告id列表"
// @Param token header string  true  "token"
// @Router /api/adverts [delete]
// @Produce json
// @Success 200 {object} response.Response{data=string}
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
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
