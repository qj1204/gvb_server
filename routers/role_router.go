package routers

import (
	"gvb_server/api"
)

func (router RouterGroup) RoleRouter() {
	app := api.ApiGroupApp.RoleApi
	router.GET("role_ids", app.RoleIDListView)
}
