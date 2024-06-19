package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	log_stash "gvb_server/plugins/log_stash_v2"
	"gvb_server/plugins/qq"
	"gvb_server/utils"
	"gvb_server/utils/jwts"
	"gvb_server/utils/pwd"
	"gvb_server/utils/random"
)

// QQLoginView qq登录，返回token，用户信息需要从token中解码
// @Tags 用户管理
// @Summary qq登录
// @Description qq登录，返回token，用户信息需要从token中解码
// @Param code query string  true  "qq登录的code"
// @Router /api/login [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) QQLoginView(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		res.FailWithMessage("没有code", c)
		return
	}
	qqInfo, err := qq.NewQQLogin(code)
	if err != nil {
		logrus.Errorf(err.Error())
		res.FailWithMessage(err.Error(), c)
		return
	}
	openID := qqInfo.OpenID
	// 根据openID判断用户是否存在
	var user models.UserModel
	err = global.DB.Take(&user, "token = ?", openID).Error
	// 随机生成16位初始密码（可以把初始密码发给用户，然后提示用户修改密码并绑定邮箱）
	initialPwd := random.RandString(16)
	ip, addr := utils.GetAddrByGin(c)
	if err != nil {
		// 不存在，就注册
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
			global.Log.Error(err)
			res.FailWithMessage("注册失败", c)
			return
		}

		// 管理员给用qq登录的用户发送初始密码
		var admin models.UserModel
		global.DB.Take(&admin, "id = ?", 1)
		err = global.DB.Create(&models.MessageModel{
			SendUserID:       admin.ID,
			SendUserNickName: admin.NickName,
			SendUserAvatar:   admin.Avatar,
			RevUserID:        user.ID,
			RevUserNickName:  user.NickName,
			RevUserAvatar:    user.Avatar,
			IsRead:           false,
			Content:          fmt.Sprintf("你的初始密码为：%s，请尽快修改密码并绑定邮箱", initialPwd),
		}).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("初始密码消息发送失败", c)
			return
		}
	}
	// 登录操作
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: user.NickName,
		Username: user.UserName,
		Role:     int(user.Role),
		UserID:   user.ID,
		//Avatar:   user.Avatar,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("token生成失败", c)
		return
	}
	c.Request.Header.Set("token", token)
	log_stash.NewSuccessLogin(c)

	global.DB.Create(&models.LoginDataModel{
		UserID:    user.ID,
		IP:        ip,
		NickName:  user.NickName,
		Token:     token,
		Device:    "web",
		Addr:      addr,
		LoginType: ctype.SignQQ,
	})

	res.OkWithData(token, c)
}
