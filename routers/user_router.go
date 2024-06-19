package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type UserRouter struct{}

/*
var store = cookie.NewStore([]byte("1429030919"))

func (this *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.UserApiGroup
	router.Use(sessions.Sessions("sessionid", store))
	router.POST("email_login", apiGroup.EmailLoginView)
	router.POST("qq_login", apiGroup.QQLoginView)
	router.POST("user", middleware.JwtAdmin(), apiGroup.UserCreateView)
	router.GET("user", middleware.JwtAuth(), apiGroup.UserListView)
	router.PUT("user_role", middleware.JwtAdmin(), apiGroup.UserUpdataRoleView)
	router.PUT("user_password", middleware.JwtAuth(), apiGroup.UserUpdatePassword)
	router.POST("user_logout", middleware.JwtAuth(), apiGroup.UserLogoutView)
	router.DELETE("user", middleware.JwtAdmin(), apiGroup.UserRemoveView)
	router.POST("user_bind_email", middleware.JwtAuth(), apiGroup.UserBindEmailView)
}
*/

func (UserRouter) InitUserRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.UserApiGroup
	router.POST("email_login", apiGroup.EmailLoginView)
	router.POST("qq_login", apiGroup.QQLoginView)
	router.GET("qq_login_path", apiGroup.QQLoginLinkView) // qq登录的跳转地址
	router.POST("users", middleware.JwtAdmin(), apiGroup.UserCreateView)
	router.GET("users", middleware.JwtAuth(), apiGroup.UserListView)
	router.PUT("user_role", middleware.JwtAdmin(), apiGroup.UserUpdateRoleView)
	router.PUT("user_password", middleware.JwtAuth(), apiGroup.UserUpdatePassword)
	router.POST("user_logout", middleware.JwtAuth(), apiGroup.UserLogoutView)
	router.DELETE("users", middleware.JwtAdmin(), apiGroup.UserRemoveView)
	router.POST("user_bind_email", middleware.JwtAuth(), apiGroup.UserBindEmailViewRedis)
	router.GET("user_info", middleware.JwtAuth(), apiGroup.UserInfoView)
	router.PUT("user_info", middleware.JwtAuth(), apiGroup.UserUpdateNickName)
}
