package article

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/es"
	"gvb_server/utils/jwt"
)

func (this *ArticleApi) ArticleCollectRemoveView(c *gin.Context) {
	var cr models.ESIDListRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	// 传入的文章id必须是该用户收藏的文章
	var articleIDList []string
	global.DB.Model(&models.UserCollectModel{}).Select("article_id").
		Where("user_id = ? and article_id in ?", claims.UserID, cr.IDList).Scan(&articleIDList)
	if len(articleIDList) == 0 {
		response.FailWithMessage("文章不存在", c)
		return
	}
	// 删除收藏记录
	global.DB.Delete(&models.UserCollectModel{}, "user_id = ? and article_id in ?", claims.UserID, cr.IDList)

	// 更新文章收藏数
	var idList []any
	for _, id := range articleIDList {
		idList = append(idList, id)
	}
	boolSearch := elastic.NewTermsQuery("_id", idList...)
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

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			response.FailWithMessage("查询失败", c)
			return
		}
		err = es.ArticleUpdate(hit.Id, map[string]any{"collects_count": article.CollectsCount - 1})
		if err != nil {
			global.Log.Error(err)
			response.FailWithMessage("更新失败", c)
			return
		}
	}
	response.OkWithMessage(fmt.Sprintf("成功取消收藏%d篇文章", len(articleIDList)), c)
}
