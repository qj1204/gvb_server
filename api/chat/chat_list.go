package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/service/common_service"
)

// ChatListView 群聊聊天记录
// @Tags 聊天管理
// @Summary 群聊聊天记录
// @Description 群聊聊天记录
// @Param data query models.Page   false  "表示多个参数"
// @Router /api/chat_groups_records [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.ChatModel]}
func (ChatApi) ChatListView(c *gin.Context) {
	var cr models.Page
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}
	cr.Sort = "created_at desc"
	list, count, _ := common_service.CommonList(models.ChatModel{IsGroup: true}, common_service.Option{
		Page:  cr,
		Debug: true,
	})

	response.OkWithList(filter.Omit("list", list), count, c)
}
