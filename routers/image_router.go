package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type ImageRouter struct{}

func (this *ImageRouter) InitImageRouter(router *gin.RouterGroup) {
	imageApiGroup := api.ApiGroupApp.ImageApiGroup
	router.POST("/image", imageApiGroup.ImageUploadView)
	router.GET("/image", imageApiGroup.ImageListView)
	router.DELETE("/image", imageApiGroup.ImageRemoveView)
	router.PUT("/image", imageApiGroup.ImageUpdateView)
	router.GET("/image_name", imageApiGroup.ImageNameListView)
}
