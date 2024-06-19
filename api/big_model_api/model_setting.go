package big_model_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwts"
	"os"
	"path"
)

const docsPath = "uploads/docs"

type ModelSetting struct {
	config.Setting
	Help string `json:"help"`
}

// ModelSettingView 获取大模型配置
func (BigModelApi) ModelSettingView(c *gin.Context) {
	token := c.GetHeader("token")
	var roleID int
	customClaims, err := jwts.ParseToken(token)
	if err == nil && customClaims != nil {
		roleID = customClaims.Role
	}
	if roleID == models.AdminRole {
		// 判断用户是不是管理员，管理员就展示所有信息
		ms := ModelSetting{
			Setting: global.Config.BigModel.Setting,
		}
		if ms.Name != "" {
			filePath := path.Join(docsPath, fmt.Sprintf("%s.md", ms.Name))
			byteData, err := os.ReadFile(filePath)
			if err == nil {
				ms.Help = string(byteData)
			}
		}
		res.OkWithData(ms, c)
		return
	}

	res.OkWithData(ModelSetting{
		Setting: config.Setting{
			Enable: global.Config.BigModel.Setting.Enable,
			Title:  global.Config.BigModel.Setting.Title,
			Slogan: global.Config.BigModel.Setting.Slogan,
			Order:  global.Config.BigModel.Setting.Order,
		},
	}, c)
}
