package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) ChatRouter() {
	app := api.ApiGroupApp.ChatApi
	router.GET("chat_groups", app.ChatGroupView)
	router.GET("chat_groups_records", app.ChatListView)
	router.DELETE("chat_groups_records", middleware.JwtAdmin(), app.ChatRemoveView)
}
