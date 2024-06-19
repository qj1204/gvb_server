package feedback_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common_service"
	"gvb_server/utils/desens"
	"gvb_server/utils/jwts"
)

// FeedBackListView 反馈列表
// @Tags 反馈管理
// @Summary 反馈列表
// @Description 反馈列表
// @Param data query models.PageInfo true  "参数"
// @Router /api/feedback [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.FeedbackModel]}
func (FeedbackApi) FeedBackListView(c *gin.Context) {
	var cr models.PageInfo
	c.ShouldBindQuery(&cr)

	var isAdmin bool

	list, count, _ := common_service.CommonList(&models.FeedbackModel{}, common_service.Option{
		PageInfo: cr,
		Likes:    []string{"email", "content"},
	})
	// 如果是普通用户和游客，则显示邮箱的第一位及后缀
	token := c.Request.Header.Get("token")
	claims, err := jwts.ParseToken(token)
	if err == nil {
		if claims.Role == 1 {
			isAdmin = true
		}
	}

	if !isAdmin {
		for _, model := range list {
			model.Email = desens.DesensitizationEmail(model.Email)
		}
	}

	res.OkWithList(list, count, c)
}
