package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type RoleDetailResponse struct {
	models.MODEL
	Icon      string        `json:"icon"`
	Name      string        `json:"name"`
	Abstract  string        `json:"abstract"`
	Tags      []TagResponse `json:"tags"`
	ChatCount int64         `json:"chatCount"`
}

type TagResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"tittle"`
	Color string `json:"color"`
}

// RoleDetailView 角色详情
func (BigModelApi) RoleDetailView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithValidError(err, c)
		return
	}
	var arm models.BigModelRoleModel
	err = global.DB.Preload("Tags").Take(&arm, cr.ID).Error
	if err != nil {
		res.FailWithMessage("角色不存在", c)
		return
	}

	var tags = make([]TagResponse, 0)
	for _, tag := range arm.Tags {
		tags = append(tags, TagResponse{
			ID:    tag.ID,
			Title: tag.Title,
			Color: tag.Color,
		})
	}
	response := RoleDetailResponse{
		MODEL:    arm.MODEL,
		Icon:     arm.Icon,
		Name:     arm.Name,
		Abstract: arm.Abstract,
		Tags:     tags,
	}

	// 找这个角色进行了多少次对话
	global.DB.Model(models.BigModelChatModel{}).Where("role_id = ?", cr.ID).Count(&response.ChatCount)

	res.OkWithData(response, c)
}
