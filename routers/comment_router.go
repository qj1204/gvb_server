package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type CommentRouter struct{}

func (this *CommentRouter) InitCommentRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.CommentApi
	router.POST("/comment", middleware.JwtAuth(), apiGroup.CommentCreateView)
	router.GET("/comment", apiGroup.CommentListView)

}
