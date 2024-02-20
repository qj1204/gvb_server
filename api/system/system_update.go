package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/common/response"
)

func (this *SystemApi) SystemInfoUpdateView(c *gin.Context) {
	var cr SystemUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.SiteInfo = info
	case "jwt":
		var info config.Jwt
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.Jwt = info
	case "qq":
		var info config.QQ
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.QQ = info
	case "qiniu":
		var info config.QiNiu
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.QiNiu = info
	case "email":
		var info config.Email
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.Email = info
	default:
		response.FailWithMessage("未找到对应的配置", c)
		return
	}

	err = core.SetYaml()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("成功", c)
}
