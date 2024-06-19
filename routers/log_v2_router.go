package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type LogV2Router struct{}

func (LogV2Router) InitLogV2Router(router *gin.RouterGroup) {
	appGroup := api.ApiGroupApp.LogV2ApiGroup
	r := router.Group("logs/v2").Use(middleware.JwtAdmin())
	r.GET("", appGroup.LogListView)      // 日志列表
	r.GET("read", appGroup.LogReadView)  // 日志读取
	r.DELETE("", appGroup.LogRemoveView) // 日志删除
}
