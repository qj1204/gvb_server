package user

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/response"
	"gvb_server/service"
	"gvb_server/utils/jwt"
)

// UserLogoutView 用户注销
// @Tags 用户管理
// @Summary 用户注销
// @Description 用户注销
// @Param token header string  true  "token"
// @Router /api/logout [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (UserApi) UserLogoutView(c *gin.Context) {
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
