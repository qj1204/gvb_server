package article

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/service/es_service"
)

type ArticleSearchRequest struct {
	models.Page
	Tag string `json:"tag" form:"tag"`
}

// ArticleListView 文章列表
// @Tags 文章管理
// @Summary 文章列表
// @Description 文章列表
// @Param data query ArticleSearchRequest   false  "表示多个参数"
// @Param token header string  false  "token"
// @Router /api/articles [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.ArticleModel]}
func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr ArticleSearchRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}
	list, count, err := es_service.CommonList(es_service.Option{
		Page:   cr.Page,
		Fields: []string{"title", "abstract", "content"},
		Tag:    cr.Tag,
	})
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("查询失败", c)
		return
	}
	// 在list场景中，过滤掉content字段
	response.OkWithList(filter.Omit("list", list), int64(count), c)
}
