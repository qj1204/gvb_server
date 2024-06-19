package message

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/service/common_service"
)

// MessageListAllView 消息列表
// @Tags 消息管理
// @Summary 消息列表
// @Description 消息列表
// @Router /api/messages_all [get]
// @Param token header string  true  "token"
// @Param data query models.Page    false  "查询参数"
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.MessageModel]}
func (MessageApi) MessageListAllView(c *gin.Context) {
	var cr models.Page
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	messageList, count, _ := common_service.CommonList(models.MessageModel{}, common_service.Option{
		Page:  cr,
		Debug: true,
	})
	response.OkWithList(messageList, count, c)
}
