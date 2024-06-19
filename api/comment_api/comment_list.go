package comment_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/redis_service"
)

type CommentListRequest struct {
	ArticleID string `form:"id" uri:"id" json:"id"`
}

// CommentListView 文章下的评论列表
// @Tags 评论管理
// @Summary 文章下的评论列表
// @Description 文章下的评论列表
// @Param id path string  true  "id"
// @Router /api/comments/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]models.CommentModel}
func (CommentApi) CommentListView(c *gin.Context) {
	var cr CommentListRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	rootCommentList := FindArticleCommentList(cr.ArticleID)
	res.OkWithList(filter.Select("c", rootCommentList), int64(len(rootCommentList)), c)
}

func FindArticleCommentList(articleID string) (rootCommentList []*models.CommentModel) {
	// 先把文章下的根评论查出来
	global.DB.Preload("User").Find(&rootCommentList, "article_id = ? and parent_comment_id is null", articleID)
	// 遍历根评论，递归查根评论下的所有子评论
	diggInfo := redis_service.NewCommentDiggCount().GetInfo()
	for i := 0; i < len(rootCommentList); i++ {
		rootDigg := diggInfo[fmt.Sprintf("%d", rootCommentList[i].ID)]
		rootCommentList[i].DiggCount = rootCommentList[i].DiggCount + rootDigg
		GetCommentTree(rootCommentList[i])
	}
	return
}

// GetCommentTree 获取评论树
func GetCommentTree(rootComment *models.CommentModel) *models.CommentModel {
	global.DB.Preload("User").Preload("SubComments").Find(rootComment)
	// 递归获取子评论树
	for _, subComment := range rootComment.SubComments {
		GetCommentTree(subComment)
	}
	return rootComment
}
