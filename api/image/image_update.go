package image

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择文件ID"`
	Name string `json:"name" binding:"required" msg:"请输入文件名"`
}

func (this *ImageApi) ImageUpdateView(c *gin.Context) {
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
