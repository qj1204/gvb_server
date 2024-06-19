package article_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_service"
	"gvb_server/service/redis_service"
	"gvb_server/utils/jwts"
	"time"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag      string `json:"tag" form:"tag"`
	Category string `json:"category" form:"category"`
	IsUser   bool   `json:"is_user" form:"is_user"` // 根据这个参数判断是否显示我收藏的文章列表
	Date     string `json:"date" form:"date"`       // 发布时间搜索
}

// ArticleListView 文章列表
// @Tags 文章管理
// @Summary 文章列表
// @Description 文章列表
// @Param data query ArticleSearchRequest   false  "表示多个参数"
// @Param token header string  false  "token"
// @Router /api/articles [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.ArticleModel]}
func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr ArticleSearchRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	boolSearch := elastic.NewBoolQuery()

	if cr.IsUser {
		token := c.GetHeader("token")
		claims, err := jwts.ParseToken(token)
		if err == nil && !redis_service.CheckLogout(token) {
			boolSearch.Must(elastic.NewTermsQuery("user_id", claims.UserID))
		}
	}

	if cr.Date != "" {
		date, err := time.Parse("2006-01-02", cr.Date)
		if err == nil {
			boolSearch.Must(elastic.NewRangeQuery("created_at").
				Gte(date.Format("2006-01-02") + " 00:00:00").
				Lte(date.Format("2006-01-02") + " 23:59:59"))
		}
	}

	list, count, err := es_service.CommonList(es_service.Option{
		PageInfo: cr.PageInfo,
		Fields:   []string{"title", "content"},
		Tag:      cr.Tag,
		Query:    boolSearch,
		Category: cr.Category,
	})
	if err != nil {
		global.Log.Error(err)
		res.OkWithMessage("查询失败", c)
		return
	}

	res.OkWithList(filter.Omit("list", list), int64(count), c)
}
