package big_model

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/response"
)

// ModelSessionSettingView 会话配置信息
func (BigModelApi) ModelSessionSettingView(c *gin.Context) {
	response.OkWithData(global.Config.BigModel.SessionSetting, c)
	return
}
