package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// ImageRemoveView 删除图片
// @Tags 图片管理
// @Summary 删除图片
// @Description 删除图片
// @Param data body models.RemoveRequest   true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/images [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文件不存在", c)
		return
	}
	err = global.DB.Delete(&imageList).Error
	if err != nil {
		res.FailWithMessage("删除图片失败，存在关联关系", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("共删除 %d 张图片", count), c)
}
