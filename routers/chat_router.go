package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type ChatRouter struct{}

func (ChatRouter) InitChatRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.ChatApiGroup
	router.GET("chat_groups", apiGroup.ChatGroupView)
	router.GET("chat_groups_records", apiGroup.ChatListView)
	router.DELETE("chat_groups_records", middleware.JwtAdmin(), apiGroup.ChatRemoveView)
}
