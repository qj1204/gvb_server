package advert

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/common"
	"strings"
)

// AdvertListView 广告列表
// @Tags 广告管理
// @Summary 广告列表
// @Description 广告列表
// @Param data query models.Page false "查询参数"
// @Router /api/advert [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.AdvertModel]}
func (this *AdvertApi) AdvertListView(c *gin.Context) {
	var cr models.Page
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	// 判断请求头的Referer是否包含admin，如果包含，就返回所有广告；不是，就返回is_show=true的广告
	referer := c.GetHeader("Referer")
	isShow := true
	if strings.Contains(referer, "admin") {
		isShow = false
	}

	// 判断Referer是否包含admin，如果包含，就全部返回；不是，就返回is_show=true的
	adverList, count, _ := common.CommonList(models.AdvertModel{IsShow: isShow}, common.Option{
		Page:  cr,
		Debug: true,
	})

	response.OkWithList(adverList, count, c)
}
