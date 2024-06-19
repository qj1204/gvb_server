package tag

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/service/es_service"
	"slices"
)

// TagRemoveView 标签删除
// @Tags 标签管理
// @Summary 标签删除
// @Description 标签删除
// @Param data body models.RemoveRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/tags [delete]
// @Produce json
// @Success 200 {object} response.Response{}
func (TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	var tagList []models.TagModel
	count := global.DB.Find(&tagList, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("标签不存在", c)
		return
	}
	global.DB.Delete(&tagList)
	// 如果标签下有文章，删除文章的标签
	for _, tag := range tagList {
		// 获取标签下的文章
		articleIDList, err := es_service.CommonIDListByTag(tag.Title)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		for _, articleID := range articleIDList {
			// 获取文章标签
			article, _ := es_service.CommonDetail(articleID)
			// 删除文章标签
			newTags := slices.DeleteFunc(article.Tags, func(s string) bool {
				return s == tag.Title
			})
			err = es_service.ArticleUpdate(articleID, map[string]any{"tags": newTags})
			if err != nil {
				global.Log.Error(err)
				continue
			}
			global.Log.Infof("文章%s的%s标签删除成功", article.Title, tag.Title)
		}
	}
	response.OkWithMessage(fmt.Sprintf("共删除%d条标签", count), c)
}
