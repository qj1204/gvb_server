package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type NewsRouter struct{}

func (this *NewsRouter) InitNewsRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.NewsApi
	router.POST("/news", apiGroup.NewsListView)
}
