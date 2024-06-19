package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type NewsRouter struct{}

func (NewsRouter) InitNewsRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.NewsApiGroup
	router.POST("news", apiGroup.NewsListView)
}
