package tag

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/common"
)

func (this *TagApi) TagListView(c *gin.Context) {
	var cr models.Page
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	tagList, count, _ := common.CommonList(models.TagModel{}, common.Option{
		Page: cr,
	})
	// 需要展示这个标签下文章的数量
	response.OkWithList(tagList, count, c)
}
