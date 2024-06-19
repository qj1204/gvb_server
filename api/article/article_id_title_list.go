package article

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
)

type ArticleIDTitleListResponse struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// ArticleIDTitleListView 文章id-title列表
// @Tags 文章管理
// @Summary 文章id-title列表
// @Description 文章id-title列表
// @Param token header string  false  "token"
// @Router /api/article_id_title [get]
// @Produce json
// @Success 200 {object} response.Response{data=[]ArticleIDTitleListResponse}
func (ArticleApi) ArticleIDTitleListView(c *gin.Context) {
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Source(`{"_source": ["title"]}`). // 这里有个bug，加上source只能获取10条数据，不知道什么原因
		Size(10000).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("查询失败", c)
		return
	}
	var articleIDTitleList = make([]ArticleIDTitleListResponse, 0)
	for _, hit := range result.Hits.Hits {
		var model models.ArticleModel
		json.Unmarshal(hit.Source, &model)
		articleIDTitleList = append(articleIDTitleList, ArticleIDTitleListResponse{
			Value: hit.Id,
			Label: model.Title,
		})
	}
	response.OkWithData(articleIDTitleList, c)
}
