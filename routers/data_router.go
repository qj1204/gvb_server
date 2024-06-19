package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) DataRouter() {
	app := api.ApiGroupApp.DataApi
	router.GET("data_login", middleware.JwtAdmin(), app.DataLoginView)
	router.GET("data_sum", middleware.JwtAdmin(), app.DataSumView)
}
