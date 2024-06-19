package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

// ModelSessionSettingView 会话配置信息
func (BigModelApi) ModelSessionSettingView(c *gin.Context) {
	res.OkWithData(global.Config.BigModel.SessionSetting, c)
}
