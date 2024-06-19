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

// FullTextSearchView 全文搜索列表
// @Tags 文章管理
// @Summary 全文搜索列表
// @Description 全文搜索列表
// @Param data query models.Page   false  "表示多个参数"
// @Router /api/articles/text [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.FullTextModel]}
func (ArticleApi) FullTextSearchView(c *gin.Context) {
	var cr models.Page
	if err := c.ShouldBindQuery(&cr); err != nil {
		global.Log.Error(err)
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	boolQuery := elastic.NewBoolQuery()
	if cr.Key != "" {
		boolQuery.Must(elastic.NewMultiMatchQuery(cr.Key, "title", "body"))
	}

	result, err := global.ESClient.
		Search(models.FullTextModel{}.Index()).
		Query(boolQuery).
		Highlight(elastic.NewHighlight().Field("body")). // 高亮显示body字段
		Size(100).
		Do(context.Background())
	if err != nil {
		return
	}
	count := result.Hits.TotalHits.Value // 搜索到的结果总条数
	fullTextList := make([]models.FullTextModel, 0)

	for _, hit := range result.Hits.Hits {
		var fullText models.FullTextModel
		_ = json.Unmarshal(hit.Source, &fullText)
		if body, ok := hit.Highlight["body"]; ok { // 如果title字段有高亮显示
			fullText.Body = body[0]
		}
		fullTextList = append(fullTextList, fullText)
	}
	response.OkWithList(fullTextList, count, c)
}
