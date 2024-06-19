package image

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择文件ID"`
	Name string `json:"name" binding:"required" msg:"请输入文件名"`
}

// ImageUpdateView 图片更新
// @Tags 图片管理
// @Summary 图片更新
// @Description 图片更新
// @Param token header string  true  "token"
// @Param data body ImageUpdateRequest   true  "表示多个参数"
// @Router /api/images [put]
// @Produce json
// @Success 200 {object} response.Response{}
func (ImageApi) ImageUpdateView(c *gin.Context) {
	var cr ImageUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
	}

	var imageModel models.BannerModel
	count := global.DB.Take(&imageModel, cr.ID).RowsAffected
	if count == 0 {
		response.FailWithMessage("图片不存在", c)
		return
	}
	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("图片名称修改成功", c)
}
