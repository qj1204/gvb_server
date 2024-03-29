package article

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/common"
	"gvb_server/utils/jwt"
)

type CollectResponse struct {
	models.ArticleModel
	CreatedAt string `json:"created_at"`
}

// ArticleCollectListView 收藏文章列表
func (this *ArticleApi) ArticleCollectListView(c *gin.Context) {
	var cr models.Page
	if err := c.ShouldBindQuery(&cr); err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var articleIDList []any
	list, count, err := common.CommonList(&models.UserCollectModel{UserID: claims.UserID}, common.Option{
		Page:  cr,
		Debug: true,
	})
	var collectMap = make(map[string]string)
	for _, v := range list {
		articleIDList = append(articleIDList, v.ArticleID)
		collectMap[v.ArticleID] = v.CreatedAt.Format("2006-01-02 15:04:05")
	}

	// 传文章id列表，查es
	boolSearch := elastic.NewTermsQuery("_id", articleIDList...)
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("查询失败", c)
		return
	}

	var res = make([]CollectResponse, 0)

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		article.ID = hit.Id
		res = append(res, CollectResponse{
			ArticleModel: article,
			CreatedAt:    collectMap[article.ID],
		})
	}
	response.OkWithList(res, count, c)
}
