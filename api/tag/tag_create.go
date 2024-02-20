package tag

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
)

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标签内容" structs:"title"` // 广告标题
}

func (this *TagApi) TagCreateView(c *gin.Context) {
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	// 重复标签判断
	var tag models.TagModel
	count := global.DB.Take(&tag, "title=?", cr.Title).RowsAffected
	if count > 0 {
		response.FailWithMessage("该标签已存在", c)
		return
	}

	err = global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("添加标签失败", c)
		return
	}
	response.OkWithMessage("添加标签成功", c)
}
