package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type UserRouter struct{}

var store = cookie.NewStore([]byte("1429030919"))

func (this *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userApiGroup := api.ApiGroupApp.UserApiGroup
	router.Use(sessions.Sessions("sessionid", store))
	router.POST("email_login", userApiGroup.EmailLoginView)
	router.POST("login", userApiGroup.QQLoginView)
	router.POST("user", middleware.JwtAdmin(), userApiGroup.UserCreateView)
	router.GET("user", middleware.JwtAuth(), userApiGroup.UserListView)
	router.PUT("user_role", middleware.JwtAdmin(), userApiGroup.UserUpdataRoleView)
	router.PUT("user_password", middleware.JwtAuth(), userApiGroup.UserUpdatePassword)
	router.POST("user_logout", middleware.JwtAuth(), userApiGroup.UserLogoutView)
	router.DELETE("user", middleware.JwtAdmin(), userApiGroup.UserRemoveView)
	router.POST("user_bind_email", middleware.JwtAuth(), userApiGroup.UserBindEmailView)
}
