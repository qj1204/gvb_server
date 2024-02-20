package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type TagRouter struct{}

func (this *TagRouter) InitTagRouter(router *gin.RouterGroup) {
	tagApiGroup := api.ApiGroupApp.TagApiGroup
	router.POST("/tag", tagApiGroup.TagCreateView)
	router.GET("/tag", tagApiGroup.TagListView)
	router.PUT("/tag/:id", tagApiGroup.TagUpdateView)
	router.DELETE("/tag", tagApiGroup.TagRemoveView)
}
