package tag

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/es"
	"slices"
)

func (this *TagApi) TagRemoveView(c *gin.Context) {
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
		articleIDList, err := es.CommonIDListByTag(tag.Title)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		for _, articleID := range articleIDList {
			// 获取文章标签
			article, _ := es.CommonDetail(articleID)
			// 删除文章标签
			newTags := slices.DeleteFunc(article.Tags, func(s string) bool {
				return s == tag.Title
			})
			err = es.ArticleUpdate(articleID, map[string]any{"tags": newTags})
			if err != nil {
				global.Log.Error(err)
				continue
			}
			global.Log.Infof("文章%s的%s标签删除成功", article.Title, tag.Title)
		}
	}
	response.OkWithMessage(fmt.Sprintf("共删除%d条标签", count), c)
}
