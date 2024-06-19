package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type RoleRouter struct{}

func (RoleRouter) InitRoleRouter(router *gin.RouterGroup) {
	appGroup := api.ApiGroupApp.RoleApiGroup
	router.GET("role_ids", appGroup.RoleIDListView)
}
