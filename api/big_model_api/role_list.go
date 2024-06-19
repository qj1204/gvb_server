package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common_service"
)

// RoleListView 列表
func (BigModelApi) RoleListView(c *gin.Context) {
	var cr models.PageInfo
	c.ShouldBindQuery(&cr)

	list, count, _ := common_service.CommonList(models.BigModelRoleModel{}, common_service.Option{
		PageInfo: cr,
		Likes:    []string{"name"},
		Preload:  []string{"Tags"},
	})
	res.OkWithList(list, count, c)
	return
}
