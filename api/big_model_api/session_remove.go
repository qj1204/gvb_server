package big_model_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// SessionRemoveView 管理员删除会话
func (BigModelApi) SessionRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithValidError(err, c)
		return
	}

	var list []models.BigModelSessionModel
	count := global.DB.Preload("ChatList").Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("记录不存在", c)
		return
	}

	if len(list) > 0 {
		// 先把引用的记录删除
		for _, session := range list {
			global.DB.Delete(&session.ChatList)
			logrus.Infof("删除关联对话 %d 条", len(session.ChatList))
		}
		err = global.DB.Delete(&list).Error
		if err != nil {
			logrus.Error(err)
			res.FailWithMessage("删除会话失败", c)
			return
		}
		logrus.Infof("删除会话 %d 个", len(list))
	}
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个会话", count), c)
}
