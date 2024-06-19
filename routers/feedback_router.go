package routers

import (
	"gvb_server/api"
)

func (router RouterGroup) FeedbackRouter() {
	app := api.ApiGroupApp.FeedbackApi
	router.POST("feedback", app.FeedBackCreateView)
	router.GET("feedback", app.FeedBackListView)
	router.DELETE("feedback", app.FeedBackRemoveView)
}
