package tag

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/service/common_service"
)

// TagListView 标签列表
// @Tags 标签管理
// @Summary 标签列表
// @Description 标签列表
// @Param data query models.Page    false  "查询参数"
// @Router /api/tags [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.TagModel]}
func (TagApi) TagListView(c *gin.Context) {
	var cr models.Page
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	tagList, count, _ := common_service.CommonList(models.TagModel{}, common_service.Option{
		Page: cr,
	})
	// 需要展示这个标签下文章的数量
	response.OkWithList(tagList, count, c)
}
