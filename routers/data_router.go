package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type DataRouter struct{}

func (DataRouter) InitDataRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.DataApiGroup
	router.GET("data_login", middleware.JwtAdmin(), apiGroup.DataLoginView)
	router.GET("data_sum", middleware.JwtAdmin(), apiGroup.DataSumView)
}
