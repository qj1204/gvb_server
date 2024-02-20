package user

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/response"
	"gvb_server/service"
	"gvb_server/utils/jwt"
)

func (this *UserApi) UserLogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	token := c.Request.Header.Get("token")

	err := service.ServiceGroupApp.UserService.Logout(claims, token)
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("注销失败", c)
		return
	}
	response.OkWithMessage("注销成功", c)
}
