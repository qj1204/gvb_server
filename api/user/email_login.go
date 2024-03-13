package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/ctype"
	"gvb_server/models/common/response"
	"gvb_server/plugins/log_stash"
	"gvb_server/utils"
	"gvb_server/utils/jwt"
	"gvb_server/utils/pwd"
)

type EmailLoginRequest struct {
	Username string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (this *UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	log := log_stash.NewLogByGin(c)

	// 判断用户是否存在
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.Username, cr.Username).Error
	if err != nil {
		// 用户不存在
		global.Log.Error("用户不存在")
		log.Warn(fmt.Sprintf("%s 用户不存在", cr.Username))
		response.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("密码错误")
		log.Warn(fmt.Sprintf("密码错误 %s %s", cr.Username, cr.Password))
		response.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 登录成功，生成token
	token, err := jwt.GenerateToken(jwt.JwtPayLoad{
		UserID: userModel.ID,
		//Username: userModel.UserName,
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
	})
	if err != nil {
		global.Log.Error("生成token失败")
		log.Error(fmt.Sprintf("生成token失败 %s", err.Error()))
		response.FailWithMessage("生成token失败", c)
		return
	}

	ip, addr := utils.GetAddrByGin(c)
	log_stash.NewLog(ip, token).Info(fmt.Sprintf("%s 登录成功", cr.Username))

	global.DB.Create(&models.LoginDataModel{
		UserID:    userModel.ID,
		IP:        ip,
		NickName:  userModel.NickName,
		Token:     token,
		Device:    "web",
		Addr:      addr,
		LoginType: ctype.SignEmail,
	})
	response.OkWithData(token, c)
}
