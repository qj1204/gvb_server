package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type ArticleRouter struct{}

func (this *ArticleRouter) InitArticleRouter(router *gin.RouterGroup) {
	articleApiGroup := api.ApiGroupApp.ArticleApiGroup
	router.POST("/article", middleware.JwtAuth(), articleApiGroup.ArticleCreateView)

}
