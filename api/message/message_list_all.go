package message

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/common"
)

func (this *MessageApi) MessageListAllView(c *gin.Context) {
	var cr models.Page
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	messageList, count, _ := common.CommonList(models.MessageModel{}, common.Option{
		Page: cr,
	})
	response.OkWithList(messageList, count, c)
}
