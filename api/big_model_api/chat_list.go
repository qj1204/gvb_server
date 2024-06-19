package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common_service"
	"gvb_server/utils/jwts"
)

type ChatListRequest struct {
	SessionID uint `json:"sessionID" form:"sessionID" binding:"required"`
	models.PageInfo
}
type ChatListResponse struct {
	models.MODEL
	UserContent string `json:"userContent"` // 用户聊天内容
	UserAvatar  string `json:"userAvatar"`  // 用户头像
	BotContent  string `json:"botContent"`  // AI的聊天内容
	BotAvatar   string `json:"botAvatar"`   // AI的头像
	Status      bool   `json:"status"`
}

// ChatListView 对话列表
func (BigModelApi) ChatListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr ChatListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithValidError(err, c)
		return
	}
	var session models.BigModelSessionModel
	err = global.DB.Take(&session, cr.SessionID).Error
	if err != nil {
		res.FailWithMessage("会话id错误", c)
		return
	}
	if claims.Role != models.AdminRole {
		// 要去验证这个会话是不是当前用户创建的
		if claims.UserID != session.UserID {
			res.FailWithMessage("会话鉴权失败", c)
			return
		}
	}

	cr.Sort = "created_at asc"

	_list, count, _ := common_service.CommonList(models.BigModelChatModel{SessionID: cr.SessionID}, common_service.Option{
		PageInfo: cr.PageInfo,
		Preload:  []string{"RoleModel", "UserModel"},
	})
	var list = make([]ChatListResponse, 0)
	for _, model := range _list {
		list = append(list, ChatListResponse{
			MODEL:       model.MODEL,
			UserContent: model.Content,
			UserAvatar:  model.UserModel.Avatar,
			BotContent:  model.BotContent,
			BotAvatar:   model.RoleModel.Icon,
			Status:      model.Status,
		})
	}

	res.OkWithList(list, count, c)
}
