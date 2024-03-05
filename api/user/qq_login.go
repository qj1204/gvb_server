package user

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/ctype"
	"gvb_server/models/common/response"
	"gvb_server/plugins/qq"
	"gvb_server/utils/jwt"
	"gvb_server/utils/pwd"
	"gvb_server/utils/random"
)

func (this *UserApi) QQLoginView(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		response.FailWithMessage("code不能为空", c)
		return
	}
	qqInfo, err := qq.NewQQLogin(code)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	openID := qqInfo.OpenID
	// 根据openID查询用户是否存在
	var user models.UserModel
	err = global.DB.Take(&user, "token = ?", openID).Error
	// 随机生成16位初始密码
	initialPwd := random.RandString(16)
	if err != nil {
		// 用户不存在，就注册
		hashPwd := pwd.HashPwd(initialPwd)
		user = models.UserModel{
			NickName:   qqInfo.Nickname,
			UserName:   openID,
			Password:   hashPwd, // 随机生成16位密码（可以把初始密码提供给用户，然后再修改密码）
			Avatar:     qqInfo.Avatar,
			Addr:       "内网", // 根据ip算地址
			Token:      openID,
			IP:         c.ClientIP(),
			Role:       ctype.PermissionUser,
			SignStatus: ctype.QQStatus,
		}
		err = global.DB.Create(&user).Error
		if err != nil {
			global.Log.Error("注册失败", err)
			response.FailWithMessage("注册失败", c)
			return
		}
	}

	// 登录操作
	token, err := jwt.GenerateToken(jwt.JwtPayLoad{
		UserID: user.ID,
		//Username: userModel.UserName,
		NickName: user.NickName,
		Role:     int(user.Role),
	})
	if err != nil {
		global.Log.Error("生成token失败")
		response.FailWithMessage("生成token失败", c)
		return
	}
	response.OkWithData(gin.H{
		"token":       token,
		"initial_pwd": initialPwd,
	}, c)
}
