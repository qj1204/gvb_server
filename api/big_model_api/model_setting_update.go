package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

// ModelSettingUpdateView 更新大模型配置
func (BigModelApi) ModelSettingUpdateView(c *gin.Context) {
	var cr config.Setting
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
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
		res.FailWithMessage("名称错误", c)
		return
	}

	global.Config.BigModel.Setting = cr
	core.SetYaml()
	res.OkWithMessage("修改成功", c)
}
