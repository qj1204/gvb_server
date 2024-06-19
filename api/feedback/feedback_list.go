package feedback

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/service/common_service"
	"gvb_server/utils/desens"
	"gvb_server/utils/jwt"
)

// FeedBackListView 反馈列表
// @Tags 反馈管理
// @Summary 反馈列表
// @Description 反馈列表
// @Param data query models.Page true  "参数"
// @Router /api/feedback [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.FeedbackModel]}
func (FeedbackApi) FeedBackListView(c *gin.Context) {
	var cr models.Page
	if err := c.ShouldBindQuery(&cr); err != nil {
		global.Log.Error(err)
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}
	var isAdmin bool

	list, count, _ := common_service.CommonList(&models.FeedbackModel{}, common_service.Option{
		Page:  cr,
		Likes: []string{"email", "content"},
	})
	// 如果是普通用户和游客，则显示邮箱的第一位及后缀
	token := c.Request.Header.Get("token")
	claims, err := jwt.ParseToken(token)
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

	response.OkWithList(list, count, c)
}
