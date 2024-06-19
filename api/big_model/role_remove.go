package big_model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
)

func (BigModelApi) RoleRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	var list []models.BigModelRoleModel
	count := global.DB.Preload("Tags").Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("记录不存在", c)
		return
	}

	if len(list) > 0 {
		// 先把引用的记录删除
		for _, i2 := range list {
			global.DB.Model(&i2).Association("Tags").Delete(i2.Tags)
		}
		err = global.DB.Delete(&list).Error
		if err != nil {
			logrus.Error(err)
			response.FailWithMessage("删除角色失败", c)
			return
		}
		logrus.Infof("删除角色 %d 个", len(list))
	}
	response.OkWithMessage(fmt.Sprintf("共删除 %d 个角色", count), c)
}
