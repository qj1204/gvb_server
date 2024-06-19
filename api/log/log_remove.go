package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/plugins/log_stash"
)

func (LogApi) LogRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	var logList []log_stash.LogStashModel
	count := global.DB.Find(&logList, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("日志不存在", c)
		return
	}
	global.DB.Delete(&logList)
	response.OkWithMessage(fmt.Sprintf("共删除%d条日志", count), c)
}
