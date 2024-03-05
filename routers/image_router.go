package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type ImageRouter struct{}

func (this *ImageRouter) InitImageRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.ImageApiGroup
	router.POST("/image", apiGroup.ImageUploadViewMy)
	router.GET("/image", apiGroup.ImageListView)
	router.DELETE("/image", apiGroup.ImageRemoveView)
	router.PUT("/image", apiGroup.ImageUpdateView)
	router.GET("/image_name", apiGroup.ImageNameListView)
}
