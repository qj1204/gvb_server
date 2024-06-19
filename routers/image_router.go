package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type ImageRouter struct{}

func (ImageRouter) InitImageRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.ImageApiGroup
	router.GET("images", apiGroup.ImageListView)
	router.GET("image_names", apiGroup.ImageNameListView)
	router.POST("images", middleware.JwtAuth(), apiGroup.ImageUploadViewMy)
	router.POST("image", middleware.JwtAuth(), apiGroup.ImageUploadDataView)
	router.DELETE("images", middleware.JwtAdmin(), apiGroup.ImageRemoveView)
	router.PUT("images", middleware.JwtAdmin(), apiGroup.ImageUpdateView)
}
