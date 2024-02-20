package tag

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
)

func (this *TagApi) TagUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	// 判断标签是否存在
	var tag models.TagModel
	count := global.DB.Take(&tag, id).RowsAffected
	if count == 0 {
		response.FailWithMessage("标签不存在", c)
		return
	}

	// 结构体转map的第三方包structs
	maps := structs.Map(&cr)
	err = global.DB.Model(&tag).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("修改标签失败", c)
		return
	}
	response.OkWithMessage("修改标签成功", c)
}
