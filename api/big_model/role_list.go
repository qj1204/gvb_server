package big_model

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/service/common_service"
)

// RoleListView 列表
func (BigModelApi) RoleListView(c *gin.Context) {
	var cr models.Page
	c.ShouldBindQuery(&cr)

	list, count, _ := common_service.CommonList(models.BigModelRoleModel{}, common_service.Option{
		Page:    cr,
		Likes:   []string{"name"},
		Preload: []string{"Tags"},
	})
	response.OkWithList(list, count, c)
	return
}
