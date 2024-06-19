package big_model

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/response"
)

// ModelSessionSettingUpdateView 更新会话配置信息
func (BigModelApi) ModelSessionSettingUpdateView(c *gin.Context) {
	var cr config.SessionSetting
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	global.Config.BigModel.SessionSetting = cr
	core.SetYaml()
	response.OkWithMessage("修改成功", c)
	return
}
