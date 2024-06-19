package message

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/service/common_service"
)

type MessageUserRecordRequest struct {
	models.Page
	SendUserID uint `json:"sendUserID" form:"sendUserID" binding:"required"`
	RevUserID  uint `json:"revUserID" form:"revUserID" binding:"required"`
}

// MessageUserRecordView 两个用户之间的聊天记录
// @Tags 消息管理
// @Summary 两个用户之间的聊天记录
// @Description 两个用户之间的聊天记录
// @Router /api/message_users/record [get]
// @Param token header string  true  "token"
// @Param data query MessageUserRecordRequest   false  "查询参数"
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.MessageModel]}
func (MessageApi) MessageUserRecordView(c *gin.Context) {
	var cr MessageUserRecordRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	list, count, _ := common_service.CommonList(models.MessageModel{}, common_service.Option{
		Page:  cr.Page,
		Where: global.DB.Where("(send_user_id = ? and rev_user_id = ? ) or ( rev_user_id = ? and send_user_id = ? )", cr.SendUserID, cr.RevUserID, cr.SendUserID, cr.RevUserID),
	})

	response.OkWithList(list, count, c)
}
