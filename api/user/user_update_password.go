package user

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/utils/jwt"
	"gvb_server/utils/pwd"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入旧密码"` // 旧密码
	NewPwd string `json:"new_pwd" binding:"required" msg:"请输入新密码"` // 新密码
}

// UserUpdatePassword 修改登录人的密码
// @Tags 用户管理
// @Summary 修改登录人的密码
// @Description 修改登录人的密码
// @Param data body UpdatePasswordRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/user_password [put]
// @Produce json
// @Success 200 {object} response.Response{}
func (UserApi) UserUpdatePassword(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		global.Log.Error("用户不存在")
		return
	}

	// 判断旧密码与数据库中的密码是否一致
	if pwd.CheckPwd(user.Password, cr.OldPwd) {
		global.Log.Error("原密码错误")
		response.FailWithMessage("原密码错误", c)
		return
	}
	// 修改密码
	hashPwd := pwd.HashPwd(cr.NewPwd)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		global.Log.Error("密码修改失败")
		response.FailWithMessage("密码修改失败", c)
		return
	}
	response.OkWithMessage("密码修改成功", c)
}
