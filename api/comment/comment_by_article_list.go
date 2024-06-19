package comment

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
)

type CommentByArticleListRequest struct {
	models.Page
	Title string `json:"title" form:"title"`
}

type CommentByArticleListResponse struct {
	Title string `json:"title"`
	ID    string `json:"id"`
	Count int    `json:"count"`
}

// CommentByArticleListView 有评论的文章列表
// @Tags 评论管理
// @Summary 有评论的文章列表
// @Description 有评论的文章列表
// @Param id path string  true  "id"
// @Param data query CommentByArticleListRequest  true  "参数"
// @Router /api/comments/articles [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[CommentByArticleListResponse]}
func (CommentApi) CommentByArticleListView(c *gin.Context) {
	var cr CommentByArticleListRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		global.Log.Error(err)
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	var count int64
	global.DB.Model(models.CommentModel{}).Group("article_id").Count(&count)

	type T struct {
		ArticleID string
		Count     int
	}
	offset := (cr.PageNum - 1) * cr.Limit

	var _list []T
	global.DB.Model(models.CommentModel{}).Group("article_id").Order("count desc").Limit(cr.Limit).Offset(offset).
		Select("article_id", "count(id) as count").Scan(&_list)

	var articleIDMap = map[string]int{}
	var articleIDList []interface{}
	for _, t := range _list {
		articleIDMap[t.ArticleID] = t.Count
		articleIDList = append(articleIDList, t.ArticleID)
	}

	res1, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewTermsQuery("_id", articleIDList...)).
		Size(10000).
		Do(context.Background())
	if err != nil {
		response.FailWithMessage("es查询错误", c)
		return
	}

	var list = make([]CommentByArticleListResponse, 0)
	for _, hit := range res1.Hits.Hits {
		var model models.ArticleModel
		err = json.Unmarshal(hit.Source, &model)
		if err != nil {
			logrus.Error(err)
			continue
		}
		model.ID = hit.Id
		list = append(list, CommentByArticleListResponse{
			Title: model.Title,
			ID:    hit.Id,
			Count: articleIDMap[hit.Id],
		})
	}
	response.OkWithList(list, count, c)
}
