package menu

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
)

type MenuNameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

// MenuNameListView 菜单名称列表
// @Tags 菜单管理
// @Summary 菜单名称列表
// @Description 菜单名称列表
// @Router /api/menu_names [get]
// @Produce json
// @Success 200 {object} response.Response{data=[]MenuNameResponse}
func (MenuApi) MenuNameListView(c *gin.Context) {
	var MenuNameResponseList []MenuNameResponse
	global.DB.Model(&models.MenuModel{}).Select("id", "title", "path").Scan(&MenuNameResponseList)
	response.OkWithData(MenuNameResponseList, c)
}
