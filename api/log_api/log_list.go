package log_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/plugins/log_stash"
	"gvb_server/service/common_service"
)

type LogRequest struct {
	models.PageInfo
	Level log_stash.Level `form:"level"`
}

// LogListView 日志列表
func (LogApi) LogListView(c *gin.Context) {
	var cr LogRequest
	c.ShouldBindQuery(&cr)
	list, count, _ := common_service.CommonList(log_stash.LogStashModel{Level: cr.Level}, common_service.Option{
		PageInfo: cr.PageInfo,
		Debug:    true,
		Likes:    []string{"ip", "addr"},
	})
	res.OkWithList(list, count, c)
	return
}
