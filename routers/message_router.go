package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type MessageRouter struct{}

func (this *MessageRouter) InitMessageRouter(router *gin.RouterGroup) {
	messageApiGroup := api.ApiGroupApp.MessageApiGroup
	router.POST("/message", middleware.JwtAuth(), messageApiGroup.MessageCreateView)
	router.GET("/message_all", middleware.JwtAdmin(), messageApiGroup.MessageListAllView)
	router.GET("/message", middleware.JwtAuth(), messageApiGroup.MessageListView)
	router.GET("/message_record", middleware.JwtAuth(), messageApiGroup.MessageRecordView)

}
