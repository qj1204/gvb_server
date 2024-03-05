package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type AdvertRouter struct{}

func (this *AdvertRouter) InitAdvertRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.AdvertApiGroup
	router.POST("/advert", apiGroup.AdvertCreateView)
	router.GET("/advert", apiGroup.AdvertListView)
	router.PUT("/advert/:id", apiGroup.AdvertUpdateView)
	router.DELETE("/advert", apiGroup.AdvertRemoveView)
}
