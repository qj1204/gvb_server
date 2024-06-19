package role_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
)

type OptionResponse struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

func (RoleApi) RoleIDListView(c *gin.Context) {
	res.OkWithData([]OptionResponse{
		{"管理员", 1},
		{"普通用户", 2},
		{"游客", 3},
	}, c)
}
