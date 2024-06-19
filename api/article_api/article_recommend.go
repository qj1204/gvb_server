package article_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_service"
)

type ArticleRecommendRequest struct {
	Category string `json:"category" uri:"category" form:"category"`
}

// ArticleRecommendView 推荐文章列表
// @Tags 文章管理
// @Summary 推荐文章列表
// @Description 推荐文章列表
// @Param data query ArticleRecommendRequest   false  "查询参数"
// @Router /api/articles/recommend [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.ArticleModel]}
func (ArticleApi) ArticleRecommendView(c *gin.Context) {
	var cr ArticleRecommendRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	list, count, err := es_service.CommonList(es_service.Option{
		PageInfo: models.PageInfo{
			Page:  1,
			Limit: 2,
		},
		Fields:   []string{"title", "content"},
		Category: cr.Category,
	})
	if err != nil {
		global.Log.Error(err)
		res.OkWithMessage("查询失败", c)
		return
	}

	res.OkWithList(filter.Omit("list", list), int64(count), c)
}
