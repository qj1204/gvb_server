package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/response"
)

type SystemUri struct {
	Name string `uri:"name"`
}

func (this *SystemApi) SystemInfoView(c *gin.Context) {
	var cr SystemUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
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
	default:
		response.FailWithMessage("未找到对应的配置", c)
	}
}
