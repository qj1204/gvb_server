package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type SystemRouter struct{}

func (this *SystemRouter) InitSystemRouter(router *gin.RouterGroup) {
	systemApiGroup := api.ApiGroupApp.SystemApiGroup
	router.GET("/system/:name", systemApiGroup.SystemInfoView)
	router.PUT("/system/:name", systemApiGroup.SystemInfoUpdateView)
}
