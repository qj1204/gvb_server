package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	log_stash "gvb_server/plugins/log_stash_v2"
	"gvb_server/service"
	"gvb_server/utils/jwts"
)

// LogoutView 用户注销
// @Tags 用户管理
// @Summary 用户注销
// @Description 用户注销
// @Param token header string  true  "token"
// @Router /api/logout [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	log := log_stash.NewAction(c)
	log.SetRequestHeader(c)
	log.SetResponse(c)
	token := c.Request.Header.Get("token")
	err := service.ServiceApp.UserService.Logout(claims, token)

	log.Info(fmt.Sprintf("用户 %s 注销登录", claims.Username))
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}

	res.OkWithMessage("注销成功", c)
}
