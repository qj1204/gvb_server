package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"gvb_server/global"
	"gvb_server/middleware"
	"net/http"
)

type RouterGroup struct {
	SystemRouterGroup   SystemRouter
	ImageRouterGroup    ImageRouter
	AdvertRouterGroup   AdvertRouter
	MenuRouterGroup     MenuRouter
	UserRouterGroup     UserRouter
	TagRouterGroup      TagRouter
	MessageRouterGroup  MessageRouter
	ArticleRouterGroup  ArticleRouter
	CommentRouterGroup  CommentRouter
	NewsRouterGroup     NewsRouter
	ChatRouterGroup     ChatRouter
	LogRouterGroup      LogRouter
	DataRouterGroup     DataRouter
	LogV2RouterGroup    LogV2Router
	RoleRouterGroup     RoleRouter
	GaodeRouterGroup    GaodeRouter
	FeedbackRouterGroup FeedbackRouter
	BigModelRouterGroup BigModelRouter
}

var RouterGroupApp = new(RouterGroup)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.Use(middleware.LogMiddleWare())
	router.StaticFS("static", http.Dir("static"))
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	apiRouterGroup := router.Group("/api")
	{
		RouterGroupApp.SystemRouterGroup.InitSystemRouter(apiRouterGroup)
		RouterGroupApp.ImageRouterGroup.InitImageRouter(apiRouterGroup)
		RouterGroupApp.AdvertRouterGroup.InitAdvertRouter(apiRouterGroup)
		RouterGroupApp.MenuRouterGroup.InitMenuRouter(apiRouterGroup)
		RouterGroupApp.UserRouterGroup.InitUserRouter(apiRouterGroup)
		RouterGroupApp.TagRouterGroup.InitTagRouter(apiRouterGroup)
		RouterGroupApp.MessageRouterGroup.InitMessageRouter(apiRouterGroup)
		RouterGroupApp.ArticleRouterGroup.InitArticleRouter(apiRouterGroup)
		RouterGroupApp.CommentRouterGroup.InitCommentRouter(apiRouterGroup)
		RouterGroupApp.NewsRouterGroup.InitNewsRouter(apiRouterGroup)
		RouterGroupApp.ChatRouterGroup.InitChatRouter(apiRouterGroup)
		RouterGroupApp.LogRouterGroup.InitLogRouter(apiRouterGroup)
		RouterGroupApp.DataRouterGroup.InitDataRouter(apiRouterGroup)
		RouterGroupApp.LogV2RouterGroup.InitLogV2Router(apiRouterGroup)
		RouterGroupApp.RoleRouterGroup.InitRoleRouter(apiRouterGroup)
		RouterGroupApp.GaodeRouterGroup.InitGaodeRouter(apiRouterGroup)
		RouterGroupApp.FeedbackRouterGroup.InitFeedbackRouter(apiRouterGroup)
		RouterGroupApp.BigModelRouterGroup.InitBigModelRouter(apiRouterGroup)
	}
	return router
}
