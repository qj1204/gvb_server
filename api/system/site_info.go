package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/response"
)

// SiteInfoView 显示网站信息
// @Tags 系统管理
// @Summary 显示网站信息
// @Description 显示网站信息
// @Router /api/settings/site [get]
// @Produce json
// @Success 200 {object} response.Response{data=config.SiteInfo}
func (SystemApi) SiteInfoView(c *gin.Context) {
	response.OkWithData(global.Config.SiteInfo, c)
}
