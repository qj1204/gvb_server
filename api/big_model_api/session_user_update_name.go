package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwts"
)

type SessionUserUpdateNameRequest struct {
	SessionID uint   `json:"sessionID" binding:"required"` // 会话id
	Name      string `json:"name"`
}

// SessionUserUpdateNameView 用户相修改会话名称
func (BigModelApi) SessionUserUpdateNameView(c *gin.Context) {
	var cr SessionUserUpdateNameRequest
	// control request
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithValidError(err, c)
		return
	}

	// 找会话
	var session models.BigModelSessionModel
	err = global.DB.Take(&session, cr.SessionID).Error
	if err != nil {
		res.FailWithMessage("会话不存在", c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	if session.UserID != claims.UserID {
		res.FailWithMessage("会话鉴权失败", c)
		return
	}
	// 修改会话名称
	err = global.DB.Model(&session).Updates(models.BigModelSessionModel{Name: cr.Name}).Error
	if err != nil {
		res.FailWithMessage("会话名称修改失败", c)
		return
	}
	res.OkWithMessage("会话名称修改成功", c)
}
