package image

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/common"
)

// ImageListView 图片列表
// @Tags 图片管理
// @Summary 图片列表
// @Description 图片列表
// @Param data query models.Page false "查询参数"
// @Router /api/image [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.BannerModel]}
func (this *ImageApi) ImageListView(c *gin.Context) {
	var cr models.Page
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	imageList, count, err := common.CommonList(models.BannerModel{}, common.Option{
		Page:  cr,
		Debug: true,
	})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithList(imageList, count, c)
}
