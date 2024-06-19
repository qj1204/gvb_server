package message

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/utils/jwt"
)

type MessageRequest struct {
	RecvUserID uint   `json:"recv_user_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

// MessageCreateView 发布消息
// @Tags 消息管理
// @Summary 发布消息
// @Description 发布消息
// @Param data body MessageRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/messages [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (MessageApi) MessageCreateView(c *gin.Context) {
	// SendUserID 就是当前登录人的id
	var cr MessageRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	var sendUser, recvUser models.UserModel
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	err = global.DB.Take(&sendUser, claims.UserID).Error
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
		SendUserID:          sendUser.ID,
		SendUserNickName:    sendUser.NickName,
		SendUserAvatar:      sendUser.Avatar,
		ReceiveUserID:       cr.RecvUserID,
		ReceiveUserNickName: recvUser.NickName,
		ReceiveUserAvatar:   recvUser.Avatar,
		IsRead:              false,
		Content:             cr.Content,
	}).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("消息发送失败", c)
		return
	}
	response.OkWithData("消息发送成功", c)
}
