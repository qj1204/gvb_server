package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/plugins/email"
	"gvb_server/service"
	"gvb_server/utils/jwts"
	"gvb_server/utils/random"
)

type BindEmailRequestRedis struct {
	Email string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code  *string `json:"code"`
	// Password string  `json:"password"` // 绑定邮箱感觉不需要密码
}

// UserBindEmailViewRedis 用户绑定邮箱
// @Tags 用户管理
// @Summary 用户绑定邮箱
// @Description 用户绑定邮箱
// @Param data body BindEmailRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/user_bind_email_redis [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserBindEmailViewRedis(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	token := c.Request.Header.Get("token")

	// 用户绑定邮箱，第一次输入是 邮箱
	// 后台会给这个邮箱发送验证码
	var cr BindEmailRequestRedis
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	if cr.Code == nil {
		// 第一次，后台发送验证码，将验证码存入redis
		code := random.Code(4)
		err = service.ServiceApp.UserService.BindEmail(token, cr.Email, code)
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("redis错误", c)
			return
		}
		email.NewCode().Send(cr.Email, fmt.Sprintf("您的验证码是：%s，有效时间为%d分钟", code, global.Config.Redis.TTL))
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("验证码发送失败", c)
			return
		}
		res.OkWithData("验证码发送成功，请查收", c)
		return
	}

	// 第二次，用户输入 邮箱 + 验证码
	email2, code, err := service.ServiceApp.UserService.GetEmailAndCodeByToken(token)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("验证码已失效", c)
		return
	}
	// 校验邮箱
	if email2 != cr.Email {
		res.FailWithMessage("邮箱错误", c)
		return
	}
	// 校验验证码
	if code != *cr.Code {
		res.FailWithMessage("验证码错误", c)
		return
	}
	// 修改用户的邮箱
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	err = global.DB.Model(&user).Updates(map[string]any{
		"email": cr.Email,
	}).Error
	if err != nil {
		res.FailWithMessage("绑定邮箱失败", c)
		return
	}
	// 删除redis中的验证码
	err = service.ServiceApp.UserService.DelEmailAndCode(token)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("redis错误", c)
		return
	}
	// 完成绑定
	res.OkWithMessage("绑定邮箱成功", c)
}