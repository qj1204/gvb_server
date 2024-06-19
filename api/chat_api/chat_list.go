package chat_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common_service"
)

// ChatListView 群聊聊天记录
// @Tags 聊天管理
// @Summary 群聊聊天记录
// @Description 群聊聊天记录
// @Param data query models.PageInfo   false  "表示多个参数"
// @Router /api/chat_groups_records [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.ChatModel]}
func (ChatApi) ChatListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	cr.Sort = "created_at desc"
	list, count, _ := common_service.CommonList(models.ChatModel{IsGroup: true}, common_service.Option{
		PageInfo: cr,
	})

	data := filter.Omit("list", list)
	res.OkWithList(data.(filter.Filter), count, c)
}
