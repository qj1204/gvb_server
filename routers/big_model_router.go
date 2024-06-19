package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type BigModelRouter struct{}

func (this *BigModelRouter) InitBigModelRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.BigModelApiGroup
	// 配置相关
	{
		router.GET("big_model/usable", middleware.JwtAdmin(), apiGroup.ModelUsableListView)                    // 可用的大模型列表
		router.GET("big_model/setting", apiGroup.ModelSettingView)                                             // 获取大模型配置
		router.PUT("big_model/setting", middleware.JwtAdmin(), apiGroup.ModelSettingUpdateView)                // 更新大模型配置
		router.GET("big_model/session_setting", middleware.JwtAdmin(), apiGroup.ModelSessionSettingView)       // 获取大模型会话配置
		router.PUT("big_model/session_setting", middleware.JwtAdmin(), apiGroup.ModelSessionSettingUpdateView) // 更新大模型会话配置
		router.PUT("big_model/auto_reply", middleware.JwtAdmin(), apiGroup.AutoReplyUpdateView)                // 自动回复添加和更新
		router.GET("big_model/auto_reply", middleware.JwtAdmin(), apiGroup.AutoReplyListView)                  // 自动回复列表
		router.DELETE("big_model/auto_reply", middleware.JwtAdmin(), apiGroup.AutoReplyRemoveView)             // 自动回复删除
	}

	// 角色相关
	{
		router.PUT("big_model/tags", middleware.JwtAdmin(), apiGroup.TagUpdateView)              // 角色标签新增和更新
		router.GET("big_model/tags", middleware.JwtAdmin(), apiGroup.TagListView)                // 角色标签分页列表
		router.GET("big_model/tags/options", middleware.JwtAdmin(), apiGroup.TagOptionsListView) // 角色标签id列表
		router.DELETE("big_model/tags", middleware.JwtAdmin(), apiGroup.TagRemoveView)           // 角色标签删除

		router.POST("big_model/roles", middleware.JwtAdmin(), apiGroup.RoleCreateView)                // 角色添加
		router.PUT("big_model/roles", middleware.JwtAdmin(), apiGroup.RoleUpdateView)                 // 角色更新
		router.GET("big_model/roles", middleware.JwtAdmin(), apiGroup.RoleListView)                   // 角色列表
		router.GET("big_model/roles_history", middleware.JwtAuth(), apiGroup.RoleUserHistoryListView) // 角色用户历史列表
		router.GET("big_model/roles/:id", apiGroup.RoleDetailView)                                    // 角色详情
		router.DELETE("big_model/roles", middleware.JwtAdmin(), apiGroup.RoleRemoveView)              // 角色删除
		router.GET("big_model/square", apiGroup.TagRoleListView)                                      // 角色广场

		router.GET("big_model/role_sessions", middleware.JwtAuth(), apiGroup.RoleSessionsView) // 角色会话列表
	}

	{
		router.GET("big_model/icons/options", apiGroup.IconsView) // 角色可选的图标
	}

	// 会话相关
	{
		router.POST("big_model/session", middleware.JwtAuth(), apiGroup.SessionCreateView)           // 用户创建会话
		router.GET("big_model/session", middleware.JwtAdmin(), apiGroup.SessionListView)             // 会话列表
		router.PUT("big_model/session", middleware.JwtAuth(), apiGroup.SessionUserUpdateNameView)    // 修改会话名称
		router.DELETE("big_model/session/:id", middleware.JwtAuth(), apiGroup.SessionUserRemoveView) // 用户删除会话
		router.DELETE("big_model/session", middleware.JwtAdmin(), apiGroup.SessionRemoveView)        // 管理员删除会话
	}
	// 对话相关
	{
		router.GET("big_model/chat_sse", apiGroup.ChatCreateView)                              // 用户创建对话
		router.GET("big_model/chat", middleware.JwtAuth(), apiGroup.ChatListView)              // 对话列表
		router.DELETE("big_model/chat/:id", middleware.JwtAuth(), apiGroup.ChatUserRemoveView) // 用户删除对话
		router.DELETE("big_model/chat", middleware.JwtAdmin(), apiGroup.ChatRemoveView)        // 管理员删除对话
	}

	// 用户相关
	{
		router.GET("big_model/user_scope_enable", middleware.JwtAuth(), apiGroup.UserScopeEnableView) // 用户是否可以领取积分
		router.POST("big_model/user_scope", middleware.JwtAuth(), apiGroup.UserScopeView)             // 用户领取积分
	}

}
