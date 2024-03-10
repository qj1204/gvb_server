package comment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/redis"
	"slices"
)

func (this *CommentApi) CommentRemoveView(c *gin.Context) {
	var cr CommentIDRequest
	if err := c.ShouldBindUri(&cr); err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	// 判断评论是否存在
	var comment models.CommentModel
	err := global.DB.Take(&comment, cr.ID).Error
	if err != nil {
		response.FailWithMessage("评论不存在", c)
		return
	}

	// 统计评论下的子评论数，再把自己算进去
	var subCommentList []models.CommentModel
	FindSubCommentCount(comment, &subCommentList)
	count := len(subCommentList) + 1
	redis.NewArticleCommentCount().SetCount(comment.ArticleID, -count)

	// 判断是否是子评论
	if comment.ParentCommentID != nil { // 是子评论
		// 找父评论，减掉对应的评论数
		global.DB.Model(&models.CommentModel{}).Where("id = ?", *comment.ParentCommentID).
			Update("comment_count", gorm.Expr("comment_count - ?", count))
	}

	// 删除子评论和根评论
	var deleteCommentIDList []uint
	for _, subCommnet := range subCommentList {
		deleteCommentIDList = append(deleteCommentIDList, subCommnet.ID)
	}
	// 翻转数组，先删除子评论，然后一个一个删
	slices.Reverse(deleteCommentIDList)
	deleteCommentIDList = append(deleteCommentIDList, comment.ID)
	for _, id := range deleteCommentIDList {
		global.DB.Delete(&models.CommentModel{}, id)
		// TODO: 删除redis中的评论点赞数
	}
	response.OkWithMessage(fmt.Sprintf("共删除%d条评论", len(deleteCommentIDList)), c)
}

func FindSubCommentCount(rootComment models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments").Take(&rootComment)
	for i := 0; i < len(rootComment.SubComments); i++ {
		*subCommentList = append(*subCommentList, rootComment.SubComments[i])
		FindSubComment(rootComment.SubComments[i], subCommentList)
	}
}
