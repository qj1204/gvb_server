package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type AdvertRouter struct{}

func (this *AdvertRouter) InitAdvertRouter(router *gin.RouterGroup) {
	advertApiGroup := api.ApiGroupApp.AdvertApiGroup
	router.POST("/advert", advertApiGroup.AdvertCreateView)
	router.GET("/advert", advertApiGroup.AdvertListView)
	router.PUT("/advert/:id", advertApiGroup.AdvertUpdateView)
	router.DELETE("/advert", advertApiGroup.AdvertRemoveView)
}
