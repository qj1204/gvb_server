package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type DiggRouter struct{}

func (this *DiggRouter) InitDiggRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.DiggApiGroup
	router.POST("/digg/article", apiGroup.DiggArticleView)

}
