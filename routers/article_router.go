package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type ArticleRouter struct{}

func (this *ArticleRouter) InitArticleRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.ArticleApiGroup
	router.POST("/article", middleware.JwtAuth(), apiGroup.ArticleCreateView)
	router.GET("/article", apiGroup.ArticleListView)
	router.GET("/article/:id", apiGroup.ArticleDetailView)
	router.GET("/article/detail", apiGroup.ArticleDetailByTitleView)
	router.GET("/article/calendar", apiGroup.ArticleCalendarView)
	router.GET("/article/tag", apiGroup.ArticleTagListView)
	router.PUT("/article", middleware.JwtAdmin(), apiGroup.ArticleUpdateView)
}
