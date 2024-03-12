package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type DataRouter struct{}

func (this *DataRouter) InitDataRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.DataApi
	router.GET("/seven_login", middleware.JwtAdmin(), apiGroup.SevenLoginView)
	router.GET("/data_sum", middleware.JwtAdmin(), apiGroup.DataSumView)
}
