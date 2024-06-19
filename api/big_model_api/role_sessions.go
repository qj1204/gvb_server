package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common_service"
	"gvb_server/utils/jwts"
	"time"
)

type RoleSessionsRequest struct {
	models.PageInfo
	RoleID uint `json:"roleID" form:"roleID" binding:"required"`
}

type RoleSessionResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

// RoleSessionsView 角色会话列表
func (BigModelApi) RoleSessionsView(c *gin.Context) {
	var cr RoleSessionsRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithValidError(err, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	_list, count, _ := common_service.CommonList(models.BigModelSessionModel{UserID: claims.UserID, RoleID: cr.RoleID}, common_service.Option{
		PageInfo: cr.PageInfo,
		Likes:    []string{"name"},
	})
	var list = make([]RoleSessionResponse, 0)
	for _, model := range _list {
		list = append(list, RoleSessionResponse{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			Name:      model.Name,
		})
	}
	res.OkWithList(list, count, c)
}
