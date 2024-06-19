package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/response"
)

// SystemInfoUpdateView 修改某一项的配置信息
// @Tags 系统管理
// @Summary 修改某一项的配置信息
// @Description 修改某一项的配置信息
// @Param name path int  true  "name"
// @Router /api/settings/{name} [put]
// @Param token header string  true  "token"
// @Produce json
// @Success 200 {object} response.Response{}
func (SystemApi) SystemInfoUpdateView(c *gin.Context) {
	var cr SystemUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}
	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(gin.ErrorTypeBind, c)
			return
		}
		global.Config.SiteInfo = info
	case "jwt":
		var info config.Jwt
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(gin.ErrorTypeBind, c)
			return
		}
		global.Config.Jwt = info
	case "qq":
		var info config.QQ
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(gin.ErrorTypeBind, c)
			return
		}
		global.Config.QQ = info
	case "qiniu":
		var info config.QiNiu
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(gin.ErrorTypeBind, c)
			return
		}
		global.Config.QiNiu = info
	case "email":
		var info config.Email
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(gin.ErrorTypeBind, c)
			return
		}
		global.Config.Email = info
	case "redis":
		var info config.Redis
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(gin.ErrorTypeBind, c)
			return
		}
		global.Config.Redis = info
	case "es":
		var info config.ES
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(gin.ErrorTypeBind, c)
			return
		}
		global.Config.ES = info
	case "chat_group":
		var info config.ChatGroup
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(gin.ErrorTypeBind, c)
			return
		}
		global.Config.ChatGroup = info
	case "gaode":
		var info config.Gaode
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(gin.ErrorTypeBind, c)
			return
		}
		if info.Key == "0d30676945160341fb0d614ef08d51ba" {
			info.Key = global.Config.Gaode.Key
		}
		global.Config.Gaode = info
	default:
		response.FailWithMessage("未找到对应的配置", c)
		return
	}

	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("配置文件修改失败", c)
		return
	}
	global.Log.Info("配置文件修改成功")
	response.OkWithMessage("成功", c)
}
