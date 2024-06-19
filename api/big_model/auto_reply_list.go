package big_model

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/service/common_service"
)

// AutoReplyListView 列表
func (BigModelApi) AutoReplyListView(c *gin.Context) {
	var cr models.Page
	c.ShouldBindQuery(&cr)

	list, count, _ := common_service.CommonList(models.AutoReplyModel{}, common_service.Option{
		Page:  cr,
		Likes: []string{"name"},
	})
	response.OkWithList(list, count, c)
	return
}
