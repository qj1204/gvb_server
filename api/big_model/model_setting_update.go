package big_model

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/response"
)

// ModelSettingUpdateView 更新大模型配置
func (BigModelApi) ModelSettingUpdateView(c *gin.Context) {
	var cr config.Setting
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	// 验证这个name有没有乱写
	var ok bool
	for _, option := range global.Config.BigModel.ModelList {
		if option.Value == cr.Name {
			ok = true
			break
		}
	}
	if !ok {
		// 没有找到
		response.FailWithMessage("名称错误", c)
		return
	}

	global.Config.BigModel.Setting = cr
	core.SetYaml()
	response.OkWithMessage("修改成功", c)
	return
}
