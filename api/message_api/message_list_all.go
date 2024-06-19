package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common_service"
)

// MessageListAllView 消息列表
// @Tags 消息管理
// @Summary 消息列表
// @Description 消息列表
// @Router /api/messages_all [get]
// @Param token header string  true  "token"
// @Param data query models.PageInfo    false  "查询参数"
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.MessageModel]}
func (MessageApi) MessageListAllView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, _ := common_service.CommonList(models.MessageModel{}, common_service.Option{
		PageInfo: cr,
	})
	res.OkWithList(list, count, c)
}
