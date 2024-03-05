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

func (this *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	apiGroup := api.ApiGroupApp.UserApiGroup
	router.POST("email_login", apiGroup.EmailLoginView)
	router.POST("qq_login", apiGroup.QQLoginView)
	router.POST("user", middleware.JwtAdmin(), apiGroup.UserCreateView)
	router.GET("user", middleware.JwtAuth(), apiGroup.UserListView)
	router.PUT("user_role", middleware.JwtAdmin(), apiGroup.UserUpdataRoleView)
	router.PUT("user_password", middleware.JwtAuth(), apiGroup.UserUpdatePassword)
	router.POST("user_logout", middleware.JwtAuth(), apiGroup.UserLogoutView)
	router.DELETE("user", middleware.JwtAdmin(), apiGroup.UserRemoveView)
	router.POST("user_bind_email", middleware.JwtAuth(), apiGroup.UserBindEmailViewRedis)
}
