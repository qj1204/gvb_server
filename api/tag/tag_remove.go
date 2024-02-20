package tag

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
)

func (this *TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var tagList []models.TagModel
	count := global.DB.Find(&tagList, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("标签不存在", c)
		return
	}
	// 如果这个标签下有文章，该怎么办？
	global.DB.Delete(&tagList)
	response.OkWithMessage(fmt.Sprintf("共删除%d条标签", count), c)
}
