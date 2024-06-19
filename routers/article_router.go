package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAdmin(), app.ArticleCreateView)
	router.GET("articles", app.ArticleListView)
	router.GET("article_id_title", app.ArticleIDTitleListView)
	router.GET("categorys", app.ArticleCategoryListView)
	router.GET("articles/detail", app.ArticleDetailByTitleView)
	router.GET("articles/calendar", app.ArticleCalendarView)
	router.GET("articles/tags", app.ArticleTagListView)
	router.PUT("articles", middleware.JwtAdmin(), app.ArticleUpdateView)
	router.DELETE("articles", middleware.JwtAdmin(), app.ArticleRemoveView)
	router.POST("articles/collects", middleware.JwtAuth(), app.ArticleCollCreateView)
	router.GET("articles/collects", middleware.JwtAuth(), app.ArticleCollListView) // 用户收藏的文章列表
	router.DELETE("articles/collects", middleware.JwtAuth(), app.ArticleCollBatchRemoveView)
	router.GET("articles/:id", app.ArticleDetailByIDView)
	router.GET("articles/recommend", app.ArticleRecommendView)
	router.POST("articles/digg", middleware.JwtAuth(), app.ArticleDiggView) // 文章点赞
	router.GET("articles/content/:id", app.ArticleContentByIDView)          // 文章内容
	router.GET("articles/text", app.FullTextSearchView)                     // 全文搜索
}
