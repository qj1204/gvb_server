package big_model

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
)

// TagOptionsListView 标签id列表
func (BigModelApi) TagOptionsListView(c *gin.Context) {
	var list []models.Options[uint]
	global.DB.Model(models.BigModelTagModel{}).Select("id as value", "title as label").Scan(&list)
	response.OkWithData(list, c)
}
