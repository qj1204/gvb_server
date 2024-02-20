package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"gvb_server/global"
)

type RouterGroup struct {
	SystemRouterGroup SystemRouter
	ImageRouterGroup  ImageRouter
	AdvertRouterGroup AdvertRouter
	MenuRouterGroup   MenuRouter
	UserRouterGroup   UserRouter
	TagRouterGroup    TagRouter
	MessageRouter     MessageRouter
}

var RouterGroupApp = new(RouterGroup)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	apiRouterGroup := router.Group("/api")
	// 系统配置api
	RouterGroupApp.SystemRouterGroup.InitSystemRouter(apiRouterGroup)
	RouterGroupApp.ImageRouterGroup.InitImageRouter(apiRouterGroup)
	RouterGroupApp.AdvertRouterGroup.InitAdvertRouter(apiRouterGroup)
	RouterGroupApp.MenuRouterGroup.InitMenuRouter(apiRouterGroup)
	RouterGroupApp.UserRouterGroup.InitUserRouter(apiRouterGroup)
	RouterGroupApp.TagRouterGroup.InitTagRouter(apiRouterGroup)
	RouterGroupApp.MessageRouter.InitMessageRouter(apiRouterGroup)
	return router
}
