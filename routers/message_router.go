package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type MessageRouter struct{}

func (this *MessageRouter) InitMessageRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.MessageApiGroup
	router.POST("/message", middleware.JwtAuth(), apiGroup.MessageCreateView)
	router.GET("/message_all", middleware.JwtAdmin(), apiGroup.MessageListAllView)
	router.GET("/message", middleware.JwtAuth(), apiGroup.MessageListView)
	router.GET("/message_record", middleware.JwtAuth(), apiGroup.MessageRecordView)
}
