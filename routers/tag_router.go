package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type TagRouter struct{}

func (TagRouter) InitTagRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.TagApiGroup
	router.POST("tags", middleware.JwtAdmin(), apiGroup.TagCreateView)
	router.GET("tags", apiGroup.TagListView)
	router.GET("tag_names", apiGroup.TagNameListView)
	router.PUT("tags/:id", middleware.JwtAdmin(), apiGroup.TagUpdateView)
	router.DELETE("tags", middleware.JwtAdmin(), apiGroup.TagRemoveView)
}
