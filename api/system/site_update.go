package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/response"
)

// SiteUpdateView 编辑网站信息
// @Tags 系统管理
// @Summary 编辑网站信息
// @Description 编辑网站信息
// @Param data body config.SiteInfo true "编辑网站信息的参数"
// @Param token header string  true  "token"
// @Router /api/settings/site [put]
// @Produce json
// @Success 200 {object} response.Response{data=config.SiteInfo}
func (SystemApi) SiteUpdateView(c *gin.Context) {
	var info config.SiteInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}
	global.Config.SiteInfo = info
	core.SetYaml()
	response.OkWithMessage("网站信息更新成功", c)
}
