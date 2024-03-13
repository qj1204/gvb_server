package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type TagRouter struct{}

func (this *TagRouter) InitTagRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.TagApiGroup
	router.POST("/tag", middleware.JwtAdmin(), apiGroup.TagCreateView)
	router.GET("/tag", middleware.JwtAdmin(), apiGroup.TagListView)
	router.PUT("/tag/:id", middleware.JwtAdmin(), apiGroup.TagUpdateView)
	router.DELETE("/tag", middleware.JwtAdmin(), apiGroup.TagRemoveView)
}
