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
	"gvb_server/service/redis"
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

	//_claims, _ := c.Get("claims")
	//claims := _claims.(*jwt.CustomClaims)

	// 如果文章删除了，用户收藏这篇文章怎么办
	// 1、顺带把与这个文章关联的收藏记录删除
	// 2、用户收藏表中，增加一个字段，记录文章是否被删除
	bulkService := global.ESClient.Bulk().Index(models.ArticleModel{}.Index()).Refresh("true")
	for _, id := range cr.IDList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
		// 删除全文所搜
		go es.DeleteFullTextByArticleID(id)
		// 删除用户收藏的文章
		go global.DB.Where("article_id = ?", id).Delete(&models.UserCollectModel{})
		// 删除数据库中的文章评论，对应的评论点赞数也要删
		var commentIDList []uint
		global.DB.Model(&models.CommentModel{}).Order("created_at desc").Select("id").Find(&commentIDList, "article_id = ?", id)
		for _, commentID := range commentIDList {
			global.DB.Delete(&models.CommentModel{}, commentID)
			redis.NewCommentDiggCount().Delete(fmt.Sprintf("%d", commentID))
		}
		// 删除redis中的文章点赞数、浏览量、评论数
		redis.NewArticleDiggCount().Delete(id)
		redis.NewArticleLookCount().Delete(id)
		redis.NewArticleCommentCount().Delete(id)
	}
	res, err := bulkService.Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("删除文章失败", c)
		return
	}
	response.OkWithMessage(fmt.Sprintf("共删除%d篇文章", len(res.Succeeded())), c)
}
