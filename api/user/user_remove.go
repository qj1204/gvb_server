package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
)

func (this *UserApi) UserRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var usertList []models.UserModel
	count := global.DB.Find(&usertList, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("用户不存在", c)
		return
	}

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// TODO: 删除用户，消息表、评论表、用户收藏的文章用户发布的文章都要删除
		err = global.DB.Delete(&usertList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("删除用户失败", c)
		return
	}
	response.OkWithMessage(fmt.Sprintf("共删除%d个用户", count), c)
}
