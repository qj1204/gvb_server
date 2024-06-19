package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwts"
)

type UserScopeEnableResponse struct {
	Enable bool `json:"enable"` // 用户能不能领取
	Scope  int  `json:"scope"`  // 能领取多少积分
}

// UserScopeEnableView 用户是否可以领取积分
func (BigModelApi) UserScopeEnableView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	// 查这个用户，今天能不能领取这个积分
	var userScopeModel models.UserScopeModel
	err := global.DB.Take(&userScopeModel, "user_id = ? and to_days(created_at)=to_days(now())", userID).Error
	var response UserScopeEnableResponse
	if err == nil {
		// 查到了
		res.OkWithData(response, c)
		return
	}
	response.Enable = true
	response.Scope = global.Config.BigModel.SessionSetting.DayScope
	res.OkWithData(response, c)
}
