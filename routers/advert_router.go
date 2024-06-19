package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type AdvertRouter struct{}

func (AdvertRouter) InitAdvertRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.AdvertApiGroup
	router.POST("adverts", middleware.JwtAdmin(), apiGroup.AdvertCreateView)
	router.GET("adverts", apiGroup.AdvertListView)
	router.PUT("adverts/:id", middleware.JwtAdmin(), apiGroup.AdvertUpdateView)
	router.DELETE("adverts", middleware.JwtAdmin(), apiGroup.AdvertRemoveView)
}
