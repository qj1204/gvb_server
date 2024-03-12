package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type LogRouter struct{}

func (this *LogRouter) InitLogRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.LogApiGroup
	router.GET("log", middleware.JwtAdmin(), apiGroup.LogListView)
	router.DELETE("log", middleware.JwtAdmin(), apiGroup.LogRemoveView)
}
