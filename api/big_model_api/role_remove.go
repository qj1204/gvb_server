package big_model_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (BigModelApi) RoleRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var list []models.BigModelRoleModel
	count := global.DB.Preload("Tags").Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("记录不存在", c)
		return
	}

	if len(list) > 0 {
		// 先把引用的记录删除
		for _, role := range list {
			global.DB.Model(&role).Association("Tags").Delete(role.Tags)
		}
		err = global.DB.Delete(&list).Error
		if err != nil {
			logrus.Error(err)
			res.FailWithMessage("删除角色失败", c)
			return
		}
		logrus.Infof("删除角色 %d 个", len(list))
	}
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个角色", count), c)
}
