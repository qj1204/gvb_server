package message

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/utils/jwt"
)

type MessageRecordRequest struct {
	UserID uint `json:"user_id" binding:"required" msg:"请输入查询的用户ID"`
}

func (this *MessageApi) MessageRecordView(c *gin.Context) {
	var cr MessageRecordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var _messageList []models.MessageModel
	var messageList = make([]models.MessageModel, 0)
	global.DB.Order("created_at asc").Find(&_messageList, "send_user_id = ? and receive_user_id = ?", claims.UserID, cr.UserID)
	for _, message := range _messageList {
		if message.SendUserID == claims.UserID || message.ReceiveUserID == claims.UserID {
			messageList = append(messageList, message)
		}
	}
	// 点开消息后，将所有消息置为已读

	response.OkWithData(messageList, c)
}
