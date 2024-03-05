package image

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
)

func (this *ImageApi) ImageRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("图片不存在", c)
		return
	}
	global.DB.Delete(&imageList)
	response.OkWithMessage(fmt.Sprintf("共删除%d张图片", count), c)
}
