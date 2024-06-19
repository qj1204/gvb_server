package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type MessageRouter struct{}

func (MessageRouter) InitMessageRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.MessageApiGroup
	router.POST("messages", middleware.JwtAuth(), apiGroup.MessageCreateView)
	router.GET("messages_all", middleware.JwtAuth(), apiGroup.MessageListAllView)
	router.GET("messages", middleware.JwtAuth(), apiGroup.MessageListView)
	router.GET("messages_record", middleware.JwtAuth(), apiGroup.MessageRecordView)

	router.GET("message_users", middleware.JwtAdmin(), apiGroup.MessageUserListView)
	router.DELETE("message_users", middleware.JwtAuth(), apiGroup.MessageRecordRemoveView) // 删除聊天记录
	router.GET("message_users/user", middleware.JwtAdmin(), apiGroup.MessageUserListByUserView)
	router.GET("message_users/record", middleware.JwtAdmin(), apiGroup.MessageUserRecordView)
	router.GET("message_users/me", middleware.JwtAuth(), apiGroup.MessageUserListByMeView)
	router.GET("message_users/record/me", middleware.JwtAuth(), apiGroup.MessageUserRecordByMeView)
}
