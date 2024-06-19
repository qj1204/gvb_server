package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) LogV2Router() {
	app := api.ApiGroupApp.LogV2Api
	r := router.Group("logs/v2").Use(middleware.JwtAdmin())
	r.GET("", app.LogListView)      // 日志列表
	r.GET("read", app.LogReadView)  // 日志读取
	r.DELETE("", app.LogRemoveView) // 日志删除
}
