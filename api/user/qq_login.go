package user

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/response"
	"gvb_server/plugins/qq"
	"gvb_server/utils"
	"gvb_server/utils/jwt"
	"gvb_server/utils/pwd"
	"gvb_server/utils/random"
)

// QQLoginView qq登录，返回token，用户信息需要从token中解码
// @Tags 用户管理
// @Summary qq登录
// @Description qq登录，返回token，用户信息需要从token中解码
// @Param code query string  true  "qq登录的code"
// @Router /api/qq_login [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (UserApi) QQLoginView(c *gin.Context) {
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
	// 随机生成16位初始密码（可以把初始密码发给用户，然后提示用户修改密码并绑定邮箱）
	initialPwd := random.RandString(16)
	ip, addr := utils.GetAddrByGin(c)
	if err != nil {
		// 用户不存在，就注册
		hashPwd := pwd.HashPwd(initialPwd)
		user = models.UserModel{
			NickName:   qqInfo.Nickname,
			UserName:   openID,
			Password:   hashPwd, // 随机生成16位密码
			Avatar:     qqInfo.Avatar,
			Addr:       addr, // 根据ip算地址
			Token:      openID,
			IP:         ip,
			Role:       ctype.PermissionUser,
			SignStatus: ctype.SignQQ,
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

	global.DB.Create(&models.LoginDataModel{
		UserID:    user.ID,
		IP:        ip,
		NickName:  user.NickName,
		Token:     token,
		Device:    "web",
		Addr:      addr,
		LoginType: ctype.SignQQ,
	})
	response.OkWithData(gin.H{
		"token":       token,
		"initial_pwd": initialPwd,
	}, c)
}
