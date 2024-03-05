package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type TagRouter struct{}

func (this *TagRouter) InitTagRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.TagApiGroup
	router.POST("/tag", apiGroup.TagCreateView)
	router.GET("/tag", apiGroup.TagListView)
	router.PUT("/tag/:id", apiGroup.TagUpdateView)
	router.DELETE("/tag", apiGroup.TagRemoveView)
}
