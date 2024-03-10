package comment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/redis"
)

type CommentListRequest struct {
	ArticleID string `json:"article_id" form:"article_id"`
}

func (this *CommentApi) CommentListView(c *gin.Context) {
	var cr CommentListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}
	rootCommentList := FindArticleCommentList(cr.ArticleID)
	response.OkWithData(filter.Select("c", rootCommentList), c)
}

func FindArticleCommentList(articleId string) (rootCommentList []models.CommentModel) {
	// 先把文章下的根评论都查出来
	global.DB.Preload("User").Find(&rootCommentList, "article_id = ? and parent_comment_id is null", articleId)
	// 遍历根评论，递归查找每个根评论下的子评论
	diggInfo := redis.NewCommentDiggCount().GetInfo()
	for i := 0; i < len(rootCommentList); i++ {
		var subCommentList []models.CommentModel
		FindSubComment(rootCommentList[i], &subCommentList)
		// 从redis中获取评论点赞数
		for j := 0; j < len(subCommentList); j++ {
			subCommentList[j].DiggCount = subCommentList[j].DiggCount + diggInfo[fmt.Sprintf("%d", subCommentList[j].ID)]
		}
		rootCommentList[i].DiggCount = rootCommentList[i].DiggCount + diggInfo[fmt.Sprintf("%d", rootCommentList[i].ID)]
		rootCommentList[i].SubComments = subCommentList
	}
	return
}

// FindSubComment 递归查找子评论
func FindSubComment(rootComment models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments.User").Take(&rootComment)
	//for i := 0; i < len(rootComment.SubComments); i++ {
	//	var subCommentList_ []models.CommentModel
	//	FindSubComment(rootComment.SubComments[i], &subCommentList_)
	//	rootComment.SubComments[i].SubComments = subCommentList_
	//	*subCommentList = append(*subCommentList, rootComment.SubComments[i])
	//}

	for i := 0; i < len(rootComment.SubComments); i++ {
		*subCommentList = append(*subCommentList, rootComment.SubComments[i])
		FindSubComment(rootComment.SubComments[i], subCommentList)
	}
}
