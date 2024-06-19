package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type MenuRouter struct{}

func (MenuRouter) InitMenuRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.MenuApiGroup
	router.POST("menus", middleware.JwtAdmin(), apiGroup.MenuCreateView)
	router.GET("menus", apiGroup.MenuListView)
	router.GET("menu_names", apiGroup.MenuNameListView)
	router.GET("menus/:id", apiGroup.MenuDetailView)
	router.GET("menus/detail", apiGroup.MenuDetailByPathView)
	router.PUT("menus/:id", middleware.JwtAdmin(), apiGroup.MenuUpdateView)
	router.DELETE("menus", middleware.JwtAdmin(), apiGroup.MenuRemoveView)
}
