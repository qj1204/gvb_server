package big_model_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/utils/jwts"
)

// ChatRemoveView 管理员删除对话
func (BigModelApi) ChatRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithValidError(err, c)
		return
	}

	var _list []models.BigModelChatModel
	var list []models.BigModelChatModel
	count := global.DB.Find(&_list, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("记录不存在", c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	if len(_list) > 0 {
		for _, chat := range _list {
			// 只有管理员或者本人才能删除对话
			if claims.Role == int(ctype.PermissionAdmin) || chat.UserID == claims.UserID {
				// 将可以删除的对话添加到list中
				list = append(list, chat)
			}
		}
		// 删除对话
		err = global.DB.Delete(&list).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("删除对话失败", c)
			return
		}
	}
	global.Log.Infof("删除对话 %d 个", len(list))
	global.Log.Infof("鉴权失败 %d 个", len(_list)-len(list))

	res.OkWithMessage(fmt.Sprintf("共删除 %d 个对话", len(list)), c)
}
