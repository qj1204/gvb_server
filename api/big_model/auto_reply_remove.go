package big_model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
)

// AutoReplyRemoveView 自动回复删除
func (BigModelApi) AutoReplyRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	var list []models.AutoReplyModel
	count := global.DB.Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("记录不存在", c)
		return
	}

	if len(list) > 0 {
		global.DB.Delete(&list)
		logrus.Infof("删除自动回复记录 %d 条", len(list))
	}
	response.OkWithMessage(fmt.Sprintf("共删除 %d 个自动回复", count), c)

}
