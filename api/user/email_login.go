package user

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
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

	// 判断用户是否存在
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.Username, cr.Username).Error
	if err != nil {
		// 用户不存在
		global.Log.Warn("用户不存在")
		response.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("密码错误")
		response.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 登录成功，生成token
	token, err := jwt.GenerateToken(jwt.JwtPayLoad{
		//Username: userModel.UserName,
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		global.Log.Error("生成token失败")
		response.FailWithMessage("生成token失败", c)
		return
	}
	response.OkWithData(token, c)
}
