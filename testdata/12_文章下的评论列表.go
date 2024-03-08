package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.DB = core.InitGorm()

	// 先把文章下的根评论都查出来
	FindArticleCommentList("ZhKLDo4Beq8OFDNuzYQB")

}

func FindArticleCommentList(articleId string) {
	// 先把文章下的根评论都查出来
	var rootCommentList []models.CommentModel
	global.DB.Find(&rootCommentList, "article_id = ? and parent_comment_id is null", articleId)
	// 遍历根评论，递归查找每个根评论下的子评论
	for i := 0; i < len(rootCommentList); i++ {
		var subCommentList []models.CommentModel
		FindSubComment(rootCommentList[i], &subCommentList)
		rootCommentList[i].SubComments = subCommentList
	}
	fmt.Println(rootCommentList)
}

// FindSubComment 递归查找子评论
func FindSubComment(comment models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments").Take(&comment)
	for _, subComment := range comment.SubComments {
		*subCommentList = append(*subCommentList, subComment)
		FindSubComment(subComment, subCommentList)
	}
}
