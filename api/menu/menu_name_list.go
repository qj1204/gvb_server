package menu

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
)

type MenuNameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

func (this *MenuApi) MenuNameListView(c *gin.Context) {
	var MenuNameResponseList []MenuNameResponse
	global.DB.Model(&models.MenuModel{}).Select("id", "title", "path").Scan(&MenuNameResponseList)
	response.OkWithData(MenuNameResponseList, c)
}
