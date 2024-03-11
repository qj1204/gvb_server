package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/common"
)

func (this *ChatApi) ChatListView(c *gin.Context) {
	var cr models.Page
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}
	cr.Sort = "created_at desc"
	list, count, _ := common.CommonList(models.ChatModel{IsGroup: true}, common.Option{
		Page:  cr,
		Debug: true,
	})

	response.OkWithList(filter.Omit("list", list), count, c)
}
