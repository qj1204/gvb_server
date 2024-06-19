package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

//var store = cookie.NewStore([]byte("HyvCD89g3VDJ9646BFGEh37GFJ"))

func (router RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	//router.Use(sessions.Sessions("sessionid", store))
	router.POST("email_login", app.EmailLoginView)
	router.POST("login", app.QQLoginView)
	router.GET("qq_login_path", app.QQLoginLinkView) // qq登录的跳转地址
	router.POST("users", middleware.JwtAdmin(), app.UserCreateView)
	router.GET("users", middleware.JwtAuth(), app.UserListView)
	router.PUT("user_role", middleware.JwtAdmin(), app.UserUpdateRoleView)
	router.PUT("user_password", middleware.JwtAuth(), app.UserUpdatePassword)
	router.POST("logout", middleware.JwtAuth(), app.LogoutView)
	router.DELETE("users", middleware.JwtAdmin(), app.UserRemoveView)
	router.POST("user_bind_email", middleware.JwtAuth(), app.UserBindEmailView)
	router.POST("user_bind_email_redis", middleware.JwtAuth(), app.UserBindEmailViewRedis)
	router.GET("user_info", middleware.JwtAuth(), app.UserInfoView)
	router.PUT("user_info", middleware.JwtAuth(), app.UserUpdateNickName)
}
