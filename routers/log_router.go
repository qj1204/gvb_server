package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type LogRouter struct{}

func (LogRouter) InitLogRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.LogApiGroup
	router.GET("logs", middleware.JwtAdmin(), apiGroup.LogListView)
	router.DELETE("logs", middleware.JwtAdmin(), apiGroup.LogRemoveView)
}
