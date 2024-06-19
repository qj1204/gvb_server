package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwts"
)

type UserScopeRequest struct {
	Status bool `json:"status"`
}

// UserScopeView 用户领取积分
func (BigModelApi) UserScopeView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	var cr UserScopeRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMessage("参数错误", c)
		return
	}

	// 查这个用户，今天能不能领取这个积分
	var userScopeModel models.UserScopeModel
	err = global.DB.Take(&userScopeModel, "user_id = ? and to_days(created_at)=to_days(now())", userID).Error
	if err == nil {
		// 查到了
		res.FailWithMessage("今日已领取积分啦", c)
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, userID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	scope := global.Config.BigModel.SessionSetting.DayScope
	// 给用户加积分
	global.DB.Model(&user).Update("scope", gorm.Expr("scope + ?", scope))

	// 给用户加积分
	// 加数据
	global.DB.Create(&models.UserScopeModel{
		UserID: userID,
		Scope:  scope,
		Status: cr.Status,
	})

	res.OkWithMessage("积分领取成功", c)
}
