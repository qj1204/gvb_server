package tag

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/service/es_service"
	"slices"
)

// TagUpdateView 更新标签
// @Tags 标签管理
// @Summary 更新标签
// @Description 更新标签
// @Param data body TagRequest  true "查询参数"
// @Param token header string  true  "token"
// @Param id path int  true  "id"
// @Router /api/tags/{id} [put]
// @Produce json
// @Success 200 {object} response.Response{}
func (TagApi) TagUpdateView(c *gin.Context) {
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
	oldTagTitle := tag.Title
	// 结构体转map的第三方包structs
	maps := structs.Map(&cr)
	err = global.DB.Model(&tag).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("修改标签失败", c)
		return
	}

	// 如果标签下有文章，更新文章的标签
	articleIDList, err := es_service.CommonIDListByTag(oldTagTitle)
	if err != nil {
		global.Log.Error(err)
		return
	}
	for _, articleID := range articleIDList {
		// 获取文章标签
		article, _ := es_service.CommonDetail(articleID)
		// 更新文章标签
		index := slices.Index(article.Tags, oldTagTitle)
		article.Tags[index] = cr.Title

		err = es_service.ArticleUpdate(articleID, map[string]any{"tags": article.Tags})
		if err != nil {
			global.Log.Error(err)
			continue
		}
		global.Log.Infof("文章%s的%s标签删除成功", article.Title, oldTagTitle)
	}

	response.OkWithMessage("修改标签成功", c)
}
