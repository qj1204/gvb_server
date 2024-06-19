package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type CommentRouter struct{}

func (CommentRouter) InitCommentRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.CommentApiGroup
	router.POST("comments", middleware.JwtAuth(), apiGroup.CommentCreateView)
	router.GET("comments", apiGroup.CommentListAllView)                                       // 评论列表
	router.GET("comments/articles", middleware.JwtAdmin(), apiGroup.CommentByArticleListView) // 有评论的文章列表
	router.GET("comments/:id", apiGroup.CommentListView)                                      // 文章下的评论列表
	router.GET("comments/digg/:id", middleware.JwtAuth(), apiGroup.CommentDiggView)           // 评论点赞
	router.DELETE("comments/:id", middleware.JwtAuth(), apiGroup.CommentRemoveView)
}
