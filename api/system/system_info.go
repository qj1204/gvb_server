package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/response"
)

type SystemUri struct {
	Name string `uri:"name"`
}

// SystemInfoView 显示某一项的配置信息
// @Tags 系统管理
// @Summary 显示某一项的配置信息
// @Description 显示某一项的配置信息  site email qq qiniu jwt chat_group
// @Param name path string  true  "name"
// @Param token header string  true  "token"
// @Router /api/settings/{name} [get]
// @Produce json
// @Success 200 {object} response.Response{}
func (SystemApi) SystemInfoView(c *gin.Context) {
	var cr SystemUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}
	switch cr.Name {
	case "site":
		response.OkWithData(global.Config.SiteInfo, c)
	case "jwt":
		response.OkWithData(global.Config.Jwt, c)
	case "qq":
		response.OkWithData(global.Config.QQ, c)
	case "qiniu":
		response.OkWithData(global.Config.QiNiu, c)
	case "email":
		response.OkWithData(global.Config.Email, c)
	case "redis":
		response.OkWithData(global.Config.Redis, c)
	case "es":
		response.OkWithData(global.Config.ES, c)
	case "chat_group":
		response.OkWithData(global.Config.ChatGroup, c)
	case "gaode":
		info := global.Config.Gaode
		info.Key = "0d30676945160341fb0d614ef08d51ba"
		response.OkWithData(info, c)
	default:
		response.FailWithMessage("未找到对应的配置", c)
	}
}
