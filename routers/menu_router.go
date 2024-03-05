package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type MenuRouter struct{}

func (this *MenuRouter) InitMenuRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.MenuApiGroup
	router.POST("/menu", apiGroup.MenuCreateView)
	router.GET("/menu", apiGroup.MenuListView)
	router.GET("/menu_name", apiGroup.MenuNameListView)
	router.GET("/menu/:id", apiGroup.MenuDetailView)
	router.PUT("/menu/:id", apiGroup.MenuUpdateView)
	router.DELETE("/menu", apiGroup.MenuRemoveView)
}
