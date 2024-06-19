package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// TagRemoveView 标签删除
func (BigModelApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMessage("参数错误", c)
		return
	}
	var list []models.BigModelTagModel
	count := global.DB.Preload("Roles").Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("记录不存在", c)
		return
	}

	if len(list) > 0 {
		// 先把引用的记录删除
		for _, tag := range list {
			global.DB.Model(&tag).Association("Roles").Delete(tag.Roles)
		}
		err = global.DB.Delete(&list).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("删除标签失败", c)
			return
		}
		global.Log.Infof("删除角色标签 %d 个", len(list))
	}

	res.OkWithMessage("删除角色标签成功", c)
}
