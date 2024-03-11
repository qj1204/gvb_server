package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type ChatRouter struct{}

func (this *ChatRouter) InitChatRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.ChatApi
	router.GET("/chat_group", apiGroup.ChatGroupView)
	router.GET("/chat_group_record", apiGroup.ChatListView)
}
