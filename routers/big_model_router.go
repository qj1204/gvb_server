package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) BigModelRouter() {
	app := api.ApiGroupApp.BigModelApi
	// 配置相关
	{
		router.GET("big_model/usable", middleware.JwtAdmin(), app.ModelUsableListView)                    // 可用的大模型列表
		router.GET("big_model/setting", app.ModelSettingView)                                             // 获取大模型配置
		router.PUT("big_model/setting", middleware.JwtAdmin(), app.ModelSettingUpdateView)                // 更新大模型配置
		router.GET("big_model/session_setting", middleware.JwtAdmin(), app.ModelSessionSettingView)       // 获取大模型会话配置
		router.PUT("big_model/session_setting", middleware.JwtAdmin(), app.ModelSessionSettingUpdateView) // 更新大模型会话配置
		router.PUT("big_model/auto_reply", middleware.JwtAdmin(), app.AutoReplyUpdateView)                // 自动回复添加和更新
		router.GET("big_model/auto_reply", middleware.JwtAdmin(), app.AutoReplyListView)                  // 自动回复列表
		router.DELETE("big_model/auto_reply", middleware.JwtAdmin(), app.AutoReplyRemoveView)             // 自动回复删除
	}

	// 角色相关
	{
		router.PUT("big_model/tags", middleware.JwtAdmin(), app.TagUpdateView)              // 角色标签新增和更新
		router.GET("big_model/tags", middleware.JwtAdmin(), app.TagListView)                // 角色标签分页列表
		router.GET("big_model/tags/options", middleware.JwtAdmin(), app.TagOptionsListView) // 角色标签id列表
		router.DELETE("big_model/tags", middleware.JwtAdmin(), app.TagRemoveView)           // 角色标签删除

		router.POST("big_model/roles", middleware.JwtAdmin(), app.RoleCreateView)                // 角色添加
		router.PUT("big_model/roles", middleware.JwtAdmin(), app.RoleUpdateView)                 // 角色更新
		router.GET("big_model/roles", middleware.JwtAdmin(), app.RoleListView)                   // 角色列表
		router.GET("big_model/roles_history", middleware.JwtAuth(), app.RoleUserHistoryListView) // 用户历史角色列表
		router.GET("big_model/roles/:id", app.RoleDetailView)                                    // 角色详情
		router.DELETE("big_model/roles", middleware.JwtAdmin(), app.RoleRemoveView)              // 角色删除
		router.GET("big_model/square", app.TagRoleListView)                                      // 角色广场

		router.GET("big_model/role_sessions", middleware.JwtAuth(), app.RoleSessionsView) // 角色会话列表
	}

	{
		router.GET("big_model/icons/options", app.IconsView) // 角色可选的图标
	}

	// 会话相关
	{
		router.POST("big_model/session", middleware.JwtAuth(), app.SessionCreateView)           // 用户创建会话
		router.GET("big_model/session", middleware.JwtAdmin(), app.SessionListView)             // 会话列表
		router.PUT("big_model/session", middleware.JwtAuth(), app.SessionUserUpdateNameView)    // 修改会话名称
		router.DELETE("big_model/session/:id", middleware.JwtAuth(), app.SessionUserRemoveView) // 用户删除会话
		router.DELETE("big_model/session", middleware.JwtAdmin(), app.SessionRemoveView)        // 管理员删除会话
	}
	// 对话相关
	{
		router.GET("big_model/chat_sse", app.ChatCreateView)                 // 用户创建对话
		router.GET("big_model/chat", middleware.JwtAuth(), app.ChatListView) // 对话列表
		//router.DELETE("big_model/chat/:id", middleware.JwtAuth(), app.ChatUserRemoveView) // 用户删除对话
		router.DELETE("big_model/chat", middleware.JwtAuth(), app.ChatRemoveView) // 删除对话
	}

	// 用户相关
	{
		router.GET("big_model/user_scope_enable", middleware.JwtAuth(), app.UserScopeEnableView) // 用户是否可以领取积分
		router.POST("big_model/user_scope", middleware.JwtAuth(), app.UserScopeView)             // 用户领取积分
	}

}
