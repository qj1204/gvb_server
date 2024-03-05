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
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr MessageRecordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	var _messageList []models.MessageModel
	var messageList = make([]models.MessageModel, 0)
	global.DB.Order("created_at asc").Find(&_messageList, "send_user_id = ? or receive_user_id = ?", claims.UserID, claims.UserID)
	for _, message := range _messageList {
		if message.SendUserID == cr.UserID || message.ReceiveUserID == cr.UserID {
			// 点开消息后，将收到的消息置为已读
			if message.SendUserID == cr.UserID {
				message.IsRead = true
				global.DB.Model(&message).Update("is_read", true)
			}
			messageList = append(messageList, message)
		}
	}
	response.OkWithData(messageList, c)
}
