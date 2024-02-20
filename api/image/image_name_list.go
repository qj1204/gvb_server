package image

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

// ImageNameListView 图片名称列表
// @Tags 图片管理
// @Summary 图片名称列表
// @Description 图片名称列表
// @Router /api/image_names [get]
// @Produce json
// @Success 200 {object} response.Response{data=[]ImageResponse}
func (this *ImageApi) ImageNameListView(c *gin.Context) {
	var imageList []ImageResponse
	global.DB.Model(&models.BannerModel{}).Select("id, path, name").Scan(&imageList)
	response.OkWithData(imageList, c)
}
