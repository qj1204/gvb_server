package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/plugins/email"
	"gvb_server/utils/jwt"
	"gvb_server/utils/pwd"
	"gvb_server/utils/random"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

func (this *UserApi) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	// 用户绑定邮箱，第一次输入是 邮箱
	// 后台会给这个邮箱发送验证码
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	session := sessions.Default(c)
	if cr.Code == nil {
		// 第一次，后台发送验证码，将验证码存入session
		code := random.Code(4)
		session.Set("valid_code", code)
		err = session.Save()
		if err != nil {
			global.Log.Error(err)
			response.FailWithMessage("session错误", c)
			return
		}
		email.NewCode().Send(cr.Email, "你的验证码是："+code)
		if err != nil {
			global.Log.Error(err)
			response.FailWithMessage("验证码发送失败", c)
			return
		}
		response.OkWithData("验证码发送成功，请查收", c)
		return
	}
	// 第二次，用户输入 邮箱 + 验证码 + 密码
	code := session.Get("valid_code")
	// 校验验证码
	if code != *cr.Code {
		response.FailWithMessage("验证码错误", c)
		return
	}
	// 修改用户的邮箱
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}
	if len(cr.Password) < 6 {
		response.FailWithMessage("密码长度不能小于6位", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.Password)
	err = global.DB.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashPwd,
	}).Error
	if err != nil {
		response.FailWithMessage("绑定邮箱失败", c)
		return
	}
	// 完成绑定
	response.OkWithMessage("绑定邮箱成功", c)
}
