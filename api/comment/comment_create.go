package comment

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/service/es_service"
	"gvb_server/service/redis_service"
	"gvb_server/utils/jwt"
)

type CommentRequest struct {
	ArticleID       string `json:"article_id" binding:"required" msg:"请选择文章"`
	Content         string `json:"content" binding:"required" msg:"请输入评论内容"`
	ParentCommentID *uint  `json:"parent_comment_id"` // 父评论ID
}

// CommentCreateView 发布评论
// @Tags 评论管理
// @Summary 发布评论
// @Description 发布评论
// @Param data body CommentRequest   true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/comments [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (CommentApi) CommentCreateView(c *gin.Context) {
	var cr CommentRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	// 文章是否存在
	_, err := es_service.CommonDetail(cr.ArticleID)
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("文章不存在", c)
		return
	}

	// 判断是否是子评论
	if cr.ParentCommentID != nil { // 如果父评论id不为0，说明该条评论是子评论
		// 判断子评论与父评论是否在同一篇文章
		var parentComment models.CommentModel
		err = global.DB.Take(&parentComment, cr.ParentCommentID).Error
		if err != nil {
			global.Log.Error(err)
			response.FailWithMessage("父评论不存在", c)
			return
		}
		if parentComment.ArticleID != cr.ArticleID {
			response.FailWithMessage("父评论与文章不匹配", c)
			return
		}
		// 给父评论的子评论数+1
		err = global.DB.Model(&parentComment).Update("comment_count", gorm.Expr("comment_count + 1")).Error
	}
	// 添加评论
	global.DB.Create(&models.CommentModel{
		ArticleID:       cr.ArticleID,
		Content:         cr.Content,
		ParentCommentID: cr.ParentCommentID,
		UserID:          claims.UserID,
	})

	// 拿到文章评论数，新的评论数存到redis
	redis_service.NewArticleCommentCount().Set(cr.ArticleID)

	response.OkWithMessage("评论成功", c)
}
