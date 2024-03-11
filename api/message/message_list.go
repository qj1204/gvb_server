package message

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/utils/jwt"
	"time"
)

type Msg struct {
	SendUserID       uint      `json:"send_user_id"`
	SendUserNickName string    `json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"send_user_avatar"`
	RecvUserID       uint      `json:"recv_user_id"`
	RecvUserNickName string    `json:"recv_user_nick_name"`
	RecvUserAvatar   string    `json:"recv_user_avatar"`
	Content          string    `json:"content"`
	CreateAt         time.Time `json:"create_at"`     // 最新的消息时间
	MessageCount     int       `json:"message_count"` // 消息条数
}

func (this *MessageApi) MessageListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var msgGroup = make(map[uint]*Msg)
	var messageList []models.MessageModel
	global.DB.Order("created_at asc").
		Find(&messageList, "send_user_id = ? or receive_user_id = ?", claims.UserID, claims.UserID)

	// 消息分组
	for _, message := range messageList {
		msg := Msg{
			SendUserID:       message.SendUserID,
			SendUserNickName: message.SendUserNickName,
			SendUserAvatar:   message.SendUserAvatar,
			RecvUserID:       message.ReceiveUserID,
			RecvUserNickName: message.ReceiveUserNickName,
			RecvUserAvatar:   message.ReceiveUserAvatar,
			Content:          message.Content,
			CreateAt:         message.CreatedAt,
			MessageCount:     1,
		}
		idNum := message.SendUserID + message.ReceiveUserID // 对当前用户来说，与其他用户的id之和保证唯一
		val, ok := msgGroup[idNum]
		if !ok {
			// 不存在
			msgGroup[idNum] = &msg
			continue
		}
		msg.MessageCount = val.MessageCount + 1
		msgGroup[idNum] = &msg
	}

	var msgList []Msg
	for _, msg := range msgGroup {
		msgList = append(msgList, *msg)
	}
	response.OkWithData(msgList, c)
}
