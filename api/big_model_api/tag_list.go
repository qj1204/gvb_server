package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common_service"
)

type TagListResponse struct {
	models.MODEL
	Title     string `json:"title"`     // 名称
	Color     string `json:"color"`     // 颜色
	RoleCount int    `json:"roleCount"` // 角色个数
}

// TagListView 标签新增和更新
func (BigModelApi) TagListView(c *gin.Context) {
	var cr models.PageInfo
	c.ShouldBindQuery(&cr)
	_list, count, _ := common_service.CommonList(models.BigModelTagModel{}, common_service.Option{
		Likes:    []string{"title"},
		Preload:  []string{"Roles"},
		PageInfo: cr,
	})
	var list = make([]TagListResponse, 0)
	for _, model := range _list {
		list = append(list, TagListResponse{
			MODEL:     model.MODEL,
			Title:     model.Title,
			Color:     model.Color,
			RoleCount: len(model.Roles),
		})
	}
	res.OkWithList(list, count, c)
}
