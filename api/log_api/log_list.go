package log_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/plugins/log_stash"
	"gvb_server/service/common"
)

type LogRequest struct {
	models.Page
	Level log_stash.Level `json:"level" form:"level"`
}

func (this *LogApi) LogListView(c *gin.Context) {
	var cr LogRequest
	c.ShouldBindQuery(&cr)

	list, count, _ := common.CommonList(log_stash.LogStashModel{Level: cr.Level}, common.Option{
		Page:  cr.Page,
		Debug: true,
		Likes: []string{"ip", "addr"},
	})
	response.OkWithList(list, count, c)
}
