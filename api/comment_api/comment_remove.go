package comment_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/redis_service"
	"gvb_server/utils/jwts"
	"slices"
)

// CommentRemoveView 删除评论
// @Tags 评论管理
// @Summary 删除评论
// @Description 删除评论
// @Param token header string  true  "token"
// @Param id path int  true  "id"
// @Router /api/comments/{id} [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (CommentApi) CommentRemoveView(c *gin.Context) {
	var cr CommentIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var commentModel models.CommentModel
	err = global.DB.Take(&commentModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("评论不存在", c)
		return
	}
	// 这条评论只能由当前登录人删除，或者管理员
	if !(commentModel.UserID == claims.UserID || claims.Role == 1) {
		res.FailWithMessage("权限错误，不可删除", c)
		return
	}

	// 统计评论下的子评论数 再把自己算上去
	subCommentList := FindAllSubCommentList(commentModel)
	count := len(subCommentList) + 1
	redis_service.NewArticleCommentCount().SetCount(commentModel.ArticleID, -count)

	// 判断是否是子评论
	if commentModel.ParentCommentID != nil {
		// 子评论
		// 找父评论，减掉对应的评论数
		global.DB.Model(&models.CommentModel{}).Where("id = ?", *commentModel.ParentCommentID).
			Update("comment_count", gorm.Expr("comment_count - ?", count))
	}

	// 删除子评论以及当前评论
	var deleteCommentIDList []uint
	for _, subComment := range subCommentList {
		deleteCommentIDList = append(deleteCommentIDList, subComment.ID)
	}
	// 反转，然后一个一个删
	slices.Reverse(deleteCommentIDList)
	deleteCommentIDList = append(deleteCommentIDList, commentModel.ID)
	for _, id := range deleteCommentIDList {
		global.DB.Delete(&models.CommentModel{}, id)
		// 删除redis中的评论点赞数
		redis_service.NewCommentDiggCount().Delete(fmt.Sprintf("%d", id))
	}
	res.OkWithMessage(fmt.Sprintf("共删除 %d 条评论", len(deleteCommentIDList)), c)
}

// FindAllSubCommentList 找一个评论的所有子评论，一维化
func FindAllSubCommentList(comment models.CommentModel) (subList []models.CommentModel) {
	global.DB.Preload("SubComments").Preload("User").Take(&comment)
	for _, subComment := range comment.SubComments {
		subList = append(subList, *subComment)
		subList = append(subList, FindAllSubCommentList(*subComment)...)
	}
	return
}
