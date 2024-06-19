package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type SystemRouter struct{}

func (SystemRouter) InitSystemRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.SystemApiGroup
	router.GET("systems/:name", apiGroup.SystemInfoView)
	router.PUT("systems/:name", middleware.JwtAdmin(), apiGroup.SystemInfoUpdateView)

	router.GET("systems/site", apiGroup.SiteInfoView)
	router.PUT("systems/site", middleware.JwtAdmin(), apiGroup.SiteUpdateView)
}
