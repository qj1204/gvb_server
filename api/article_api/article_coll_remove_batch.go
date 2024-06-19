package article_api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_service"
	"gvb_server/utils/jwts"
)

// ArticleCollBatchRemoveView 用户取消收藏文章
// @Tags 文章管理
// @Summary 用户取消收藏文章
// @Description 用户取消收藏文章
// @Param data body models.ESIDListRequest   true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/articles/collects [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleCollBatchRemoveView(c *gin.Context) {
	var cr models.ESIDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var collects []models.UserCollectModel
	var articleIDList []string
	global.DB.Find(&collects, "user_id = ? and article_id in ?", claims.UserID, cr.IDList).
		Select("article_id").Scan(&articleIDList)
	if len(articleIDList) == 0 {
		res.FailWithMessage("文章不存在", c)
		return
	}

	// 更新文章收藏数
	var idList []interface{}
	for _, s := range articleIDList {
		idList = append(idList, s)
	}
	boolSearch := elastic.NewTermsQuery("_id", idList...)
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		err = es_service.ArticleUpdate(hit.Id, map[string]any{
			"collects_count": article.CollectsCount - 1,
		})
		if err != nil {
			global.Log.Error(err)
			continue
		}
	}
	global.DB.Delete(&collects)
	res.OkWithMessage(fmt.Sprintf("成功取消收藏 %d 篇文章", len(articleIDList)), c)
}
