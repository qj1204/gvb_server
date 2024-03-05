package message

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
)

type MessageRequest struct {
	SendUserID uint   `json:"send_user_id" binding:"required"` // 就是当前登录人的ID
	RecvUserID uint   `json:"recv_user_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

// MessageCreateView 发送消息
func (this *MessageApi) MessageCreateView(c *gin.Context) {
	var cr MessageRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	var sendUser, recvUser models.UserModel
	err = global.DB.Take(&sendUser, cr.SendUserID).Error
	if err != nil {
		response.FailWithMessage("发送人不存在", c)
		return
	}
	err = global.DB.Take(&recvUser, cr.RecvUserID).Error
	if err != nil {
		response.FailWithMessage("接收人不存在", c)
		return
	}
	err = global.DB.Create(&models.MessageModel{
		SendUserID:          cr.SendUserID,
		SendUserNickName:    sendUser.NickName,
		SendUserAvatar:      sendUser.Avatar,
		ReceiveUserID:       cr.RecvUserID,
		ReceiveUserNickName: recvUser.NickName,
		ReceiveUserAvatar:   recvUser.Avatar,
		IsRead:              false,
		Content:             cr.Content,
	}).Error
	if err != nil {
		response.FailWithMessage("消息发送失败", c)
		return
	}
	response.OkWithData("消息发送成功", c)
}
