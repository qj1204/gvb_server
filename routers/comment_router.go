package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type CommentRouter struct{}

func (this *CommentRouter) InitCommentRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.CommentApiGroup
	router.POST("/comment", middleware.JwtAuth(), apiGroup.CommentCreateView)
	router.GET("/comment", apiGroup.CommentListView)
	router.GET("/comment/:id", middleware.JwtAuth(), apiGroup.CommentDiggView)
	router.DELETE("/comment/:id", middleware.JwtAuth(), apiGroup.CommentRemoveView)
}
