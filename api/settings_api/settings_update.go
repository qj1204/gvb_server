package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

// SettingsInfoUpdateView 修改某一项的配置信息
// @Tags 系统管理
// @Summary 修改某一项的配置信息
// @Description 修改某一项的配置信息
// @Param name path int  true  "name"
// @Router /api/settings/{name} [put]
// @Param token header string  true  "token"
// @Produce json
// @Success 200 {object} res.Response{}
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "email":
		var info config.Email
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		// 判断是不是传的******
		if info.Password == "******" {
			info.Password = global.Config.Email.Password
		}
		global.Config.Email = info
	case "qq":
		var info config.QQ
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		if info.Key == "******" {
			info.Key = global.Config.QQ.Key
		}
		global.Config.QQ = info
	case "qiniu":
		var info config.QiNiu
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		if info.SecretKey == "******" {
			info.SecretKey = global.Config.QiNiu.SecretKey
		}
		global.Config.QiNiu = info
	case "jwt":
		var info config.Jwt
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		if info.Secret == "******" {
			info.Secret = global.Config.Jwt.Secret
		}
		global.Config.Jwt = info
	case "chat_group":
		var info config.ChatGroup
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.ChatGroup = info
	case "gaode":
		var info config.Gaode
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		if info.Key == "******" {
			info.Key = global.Config.Gaode.Key
		}
		global.Config.Gaode = info
	default:
		res.FailWithMessage("没有对应的配置信息", c)
		return
	}

	core.SetYaml()
	res.OkWith(c)
}
