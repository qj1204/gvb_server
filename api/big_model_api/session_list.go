package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common_service"
)

type SessionListRequest struct {
	models.PageInfo
}
type SessionListResponse struct {
	models.MODEL
	UserID      uint   `json:"userID"`
	NickName    string `json:"nickName"`
	SessionName string `json:"sessionName"` // 会话名称   有名称就用自己的名称，没有就自动生成
	RoleName    string `json:"roleName"`    // ai角色的名称
	ChatCount   int    `json:"chatCount"`   // 对话的次数
	LastContent string `json:"lastContent"` // 最后一次的聊天内容
}

// SessionListView 会话列表
func (BigModelApi) SessionListView(c *gin.Context) {
	var cr SessionListRequest
	c.ShouldBindQuery(&cr)
	_list, count, _ := common_service.CommonList(models.BigModelSessionModel{}, common_service.Option{
		Preload: []string{"UserModel", "RoleModel", "ChatList"},
	})
	var list = make([]SessionListResponse, 0)
	for _, model := range _list {
		var lastContent string
		if len(model.ChatList) > 0 {
			lastContent = model.ChatList[len(model.ChatList)-1].Content
		}

		list = append(list, SessionListResponse{
			MODEL:       model.MODEL,
			NickName:    model.UserModel.NickName,
			SessionName: model.Name,
			RoleName:    model.RoleModel.Name,
			ChatCount:   len(model.ChatList),
			LastContent: lastContent,
			UserID:      model.UserID,
		})
	}
	res.OkWithList(list, count, c)
}
