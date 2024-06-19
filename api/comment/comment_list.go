package comment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/service/redis_service"
)

type CommentListRequest struct {
	ArticleID string `json:"id" form:"id" uri:"id"`
}

// CommentListView 文章下的评论列表
// @Tags 评论管理
// @Summary 文章下的评论列表
// @Description 文章下的评论列表
// @Param id path string  true  "id"
// @Router /api/comments/{id} [get]
// @Produce json
// @Success 200 {object} response.Response{data=[]models.CommentModel}
func (CommentApi) CommentListView(c *gin.Context) {
	var cr CommentListRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		global.Log.Error(err)
		response.FailWithError(err, &cr, c)
		return
	}
	rootCommentList := FindArticleCommentList(cr.ArticleID)
	response.OkWithList(filter.Select("c", rootCommentList), int64(len(rootCommentList)), c)
	return
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
	fmt.Println(rootComment.SubComments)
	// 递归获取子评论树
	for _, subComment := range rootComment.SubComments {
		GetCommentTree(subComment)
	}
	return rootComment
}

//func (CommentApi) CommentListView(c *gin.Context) {
//	var cr CommentListRequest
//	err := c.ShouldBindUri(&cr)
//	if err != nil {
//		response.FailWithCode(gin.ErrorTypeBind, c)
//		return
//	}
//	rootCommentList := FindArticleCommentList(cr.ArticleID)
//	response.OkWithData(filter.Select("c", rootCommentList), c)
//}
//
//func FindArticleCommentList(articleId string) (rootCommentList []models.CommentModel) {
//	// 先把文章下的根评论都查出来
//	global.DB.Preload("User").Find(&rootCommentList, "article_id = ? and parent_comment_id is null", articleId)
//	// 遍历根评论，递归查找每个根评论下的子评论
//	diggInfo := redis_service.NewCommentDiggCount().GetInfo()
//	for i := 0; i < len(rootCommentList); i++ {
//		var subCommentList []models.CommentModel
//		FindSubComment(rootCommentList[i], &subCommentList)
//		// 从redis中获取评论点赞数
//		for j := 0; j < len(subCommentList); j++ {
//			subCommentList[j].DiggCount = subCommentList[j].DiggCount + diggInfo[fmt.Sprintf("%d", subCommentList[j].ID)]
//		}
//		rootCommentList[i].DiggCount = rootCommentList[i].DiggCount + diggInfo[fmt.Sprintf("%d", rootCommentList[i].ID)]
//		rootCommentList[i].SubComments = subCommentList
//	}
//	return
//}
//
//// FindSubComment 递归查找子评论
//func FindSubComment(rootComment models.CommentModel, subCommentList *[]models.CommentModel) {
//	global.DB.Preload("SubComments.User").Take(&rootComment)
//	//for i := 0; i < len(rootComment.SubComments); i++ {
//	//	var subCommentList_ []models.CommentModel
//	//	FindSubComment(rootComment.SubComments[i], &subCommentList_)
//	//	rootComment.SubComments[i].SubComments = subCommentList_
//	//	*subCommentList = append(*subCommentList, rootComment.SubComments[i])
//	//}
//
//	for i := 0; i < len(rootComment.SubComments); i++ {
//		*subCommentList = append(*subCommentList, rootComment.SubComments[i])
//		FindSubComment(rootComment.SubComments[i], subCommentList)
//	}
//}
