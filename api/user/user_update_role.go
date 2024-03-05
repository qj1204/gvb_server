package user

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/ctype"
	"gvb_server/models/common/response"
)

type UserRole struct {
	UserID   uint       `json:"user_id" binding:"required" msg:"用户id错误"`
	NickName string     `json:"nick_name"` // 防止用户昵称非法，管理员有能力修改
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
}

func (this *UserApi) UserUpdataRoleView(c *gin.Context) {
	var cr UserRole
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		response.FailWithMessage("用户ID错误，用户不存在", c)
		return
	}
	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("修改权限失败", c)
		return
	}
	response.OkWithMessage("修改权限成功", c)
}
