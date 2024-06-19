package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type ArticleRouter struct{}

func (ArticleRouter) InitArticleRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.ArticleApiGroup
	router.POST("articles", middleware.JwtAdmin(), apiGroup.ArticleCreateView)
	router.GET("articles", apiGroup.ArticleListView)
	router.GET("article_id_title", apiGroup.ArticleIDTitleListView)
	router.GET("articles/categorys", apiGroup.ArticleCategoryListView)
	router.GET("articles/detail", apiGroup.ArticleDetailByTitleView)
	router.GET("articles/calendar", apiGroup.ArticleCalendarView)
	router.GET("articles/tags", apiGroup.ArticleTagListView)
	router.PUT("articles", middleware.JwtAdmin(), apiGroup.ArticleUpdateView)
	router.DELETE("articles", middleware.JwtAdmin(), apiGroup.ArticleRemoveView)
	router.POST("articles/collects", middleware.JwtAuth(), apiGroup.ArticleCollectCreateView)
	router.GET("articles/collects", middleware.JwtAuth(), apiGroup.ArticleCollectListView)
	router.DELETE("articles/collects", middleware.JwtAuth(), apiGroup.ArticleCollectRemoveView)
	router.GET("articles/content/:id", apiGroup.ArticleContentByIDView) // 文章内容
	router.GET("articles/:id", apiGroup.ArticleDetailByIDView)
	router.POST("articles/digg", apiGroup.ArticleDiggView)
	router.GET("articles/text", apiGroup.FullTextSearchView)
}
