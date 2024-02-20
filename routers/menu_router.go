package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type MenuRouter struct{}

func (this *MenuRouter) InitMenuRouter(router *gin.RouterGroup) {
	menuApiGroup := api.ApiGroupApp.MenuApiGroup
	router.POST("/menu", menuApiGroup.MenuCreateView)
	router.GET("/menu", menuApiGroup.MenuListView)
	router.GET("/menu_names", menuApiGroup.MenuNameListView)
	router.GET("/menu/:id", menuApiGroup.MenuDetailView)
	router.PUT("/menu/:id", menuApiGroup.MenuUpdateView)
	router.DELETE("/menu", menuApiGroup.MenuRemoveView)
}
