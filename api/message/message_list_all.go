package message

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/common"
)

func (this *MessageApi) MessageListAllView(c *gin.Context) {
	var cr models.Page
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	messageList, count, _ := common.CommonList(models.MessageModel{}, common.Option{
		Page:  cr,
		Debug: true,
	})
	fmt.Println(count)
	response.OkWithList(messageList, count, c)
}
