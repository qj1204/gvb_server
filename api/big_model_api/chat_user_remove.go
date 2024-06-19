package big_model_api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwts"
)

// ChatUserRemoveView 用户删除对话
func (BigModelApi) ChatUserRemoveView(c *gin.Context) {
	var cr models.IDRequest
	// control request
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithValidError(err, c)
		return
	}

	// 找会话
	var chat models.BigModelChatModel
	err = global.DB.Take(&chat, cr.ID).Error
	if err != nil {
		res.FailWithMessage("对话不存在", c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	if chat.UserID != claims.UserID {
		res.FailWithMessage("对话鉴权失败", c)
		return
	}
	// 删除对话
	err = global.DB.Delete(&chat).Error
	if err != nil {
		logrus.Error(err)
		res.FailWithMessage("对话删除失败", c)
		return
	}
	res.OkWithMessage("对话删除成功", c)
}
