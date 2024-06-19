package advert

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/service/common_service"
	"strings"
)

// AdvertListView 广告列表
// @Tags 广告管理
// @Summary 广告列表
// @Description 广告列表
// @Param data query models.Page false "查询参数"
// @Router /api/adverts [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.AdvertModel]}
func (AdvertApi) AdvertListView(c *gin.Context) {
	var cr models.Page
	if err := c.ShouldBindQuery(&cr); err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	// 判断请求头的Gvb_referer是否包含admin，如果包含，就返回所有广告；不是，就返回is_show=true的广告
	referer := c.GetHeader("Gvb_referer")
	isShow := true
	if strings.Contains(referer, "admin") {
		// admin来的
		isShow = false
	}

	advertList, count, _ := common_service.CommonList(models.AdvertModel{IsShow: isShow}, common_service.Option{
		Page:  cr,
		Debug: true,
	})

	response.OkWithList(advertList, count, c)
}
