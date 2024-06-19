package big_model

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/utils/jwt"
)

type UserScopeEnableResponse struct {
	Enable bool `json:"enable"` // 用户能不能领取
	Scope  int  `json:"scope"`  // 能领取多少积分
}

// UserScopeEnableView 用户是否可以领取积分
func (BigModelApi) UserScopeEnableView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	userID := claims.UserID

	// 查这个用户，今天能不能领取这个积分
	var userScopeModel models.UserScopeModel
	err := global.DB.Take(&userScopeModel, "user_id = ? and to_days(created_at)=to_days(now())", userID).Error
	var res UserScopeEnableResponse
	if err == nil {
		// 查到了
		response.OkWithData(res, c)
		return
	}
	res.Enable = true
	res.Scope = global.Config.BigModel.SessionSetting.DayScope
	response.OkWithData(res, c)
}
