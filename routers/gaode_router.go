package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) GaodeRouter() {
	app := api.ApiGroupApp.GaodeApi
	r := router.Group("gaode")
	r.GET("weather", middleware.JwtAuth(), app.WeatherInfoView)
}
