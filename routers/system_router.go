package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type SystemRouter struct{}

func (this *SystemRouter) InitSystemRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.SystemApiGroup
	router.GET("/system/:name", apiGroup.SystemInfoView)
	router.PUT("/system/:name", apiGroup.SystemInfoUpdateView)
}
