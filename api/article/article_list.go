package article

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/es"
)

func (this *ArticleApi) ArticleListView(c *gin.Context) {
	var page models.Page
	if err := c.ShouldBindQuery(&page); err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}
	list, count, err := es.CommonList(page.Key, page.PageNum, page.Limit)
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("查询失败", c)
		return
	}
	// 在list场景中，过滤掉content字段
	response.OkWithList(filter.Omit("list", list), int64(count), c)
}
