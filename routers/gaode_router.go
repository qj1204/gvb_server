package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type GaodeRouter struct{}

func (GaodeRouter) InitGaodeRouter(router *gin.RouterGroup) {
	appGroup := api.ApiGroupApp.GaodeApiGroup
	router.GET("gaode/weather", middleware.JwtAuth(), appGroup.WeatherInfoView)
}
