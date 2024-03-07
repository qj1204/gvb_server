package article

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/es"
	"gvb_server/utils/jwt"
)

type IDListRequest struct {
	IDList []string `json:"id_list"`
}

func (this *ArticleApi) ArticleRemoveView(c *gin.Context) {
	var cr IDListRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	// 如果文章删除了，用户收藏这篇文章怎么办
	// 1、顺带把与这个文章关联的收藏记录删除
	// 2、用户收藏表中，增加一个字段，记录文章是否被删除
	bulkService := global.ESClient.Bulk().Index(models.ArticleModel{}.Index()).Refresh("true")
	for _, id := range cr.IDList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
		go es.DeleteFullTextByArticleID(id)
		go global.DB.Where("user_id = ? and article_id = ?", claims.UserID, id).Delete(&models.UserCollectModel{})
	}
	res, err := bulkService.Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("删除文章失败", c)
		return
	}
	response.OkWithMessage(fmt.Sprintf("共删除%d篇文章", len(res.Succeeded())), c)
}
