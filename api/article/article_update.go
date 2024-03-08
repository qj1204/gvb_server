package article

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/ctype"
	"gvb_server/models/common/response"
	"gvb_server/service/es"
	"time"
)

type ArticleUpdateRequest struct {
	ID       string      `json:"id"`
	Title    string      `json:"title"`     // 文章标题
	Abstract string      `json:"abstract"`  // 文章简介
	Content  string      `json:"content"`   // 文章内容
	Category string      `json:"category"`  // 文章分类
	Source   string      `json:"source"`    // 文章来源
	Link     string      `json:"link"`      // 原文链接
	BannerID uint        `json:"banner_id"` // 文章封面ID
	Tags     ctype.Array `json:"tags"`      // 文章标签
}

func (this *ArticleApi) ArticleUpdateView(c *gin.Context) {
	var cr ArticleUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		response.FailWithError(err, &cr, c)
		return
	}

	// 判断文章是否存在
	oldArticle, err := es.CommonDetail(cr.ID)
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("文章不存在", c)
		return
	}

	// 更新banner
	var bannerUrl string
	if cr.BannerID != 0 {
		err = global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl).Error
		if err != nil {
			global.Log.Error(err)
			response.FailWithMessage("banner不存在", c)
			return
		}
	}

	tmpArticle := models.ArticleModel{
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Title:     cr.Title,
		Keyword:   cr.Title,
		Abstract:  cr.Abstract,
		Content:   cr.Content,
		Category:  cr.Category,
		Source:    cr.Source,
		Link:      cr.Link,
		BannerID:  cr.BannerID,
		BannerUrl: bannerUrl,
		Tags:      cr.Tags,
	}

	maps := structs.Map(&tmpArticle)
	// 去掉空值
	for k, v := range maps {
		switch val := v.(type) {
		case string:
			if val == "" {
				delete(maps, k)
			}
		case uint:
			if val == 0 {
				delete(maps, k)
			}
		case int:
			if val == 0 {
				delete(maps, k)
			}
		case ctype.Array:
			if len(val) == 0 {
				delete(maps, k)
			}
		}
	}

	err = es.ArticleUpdate(cr.ID, maps)
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("文章更新失败", c)
		return
	}

	// 更新全文搜索
	newArticle, _ := es.CommonDetail(cr.ID)
	if oldArticle.Title != newArticle.Title || oldArticle.Content != newArticle.Content {
		go es.DeleteFullTextByArticleID(cr.ID)
		go es.AsyncArticleByFullText(cr.ID, newArticle.Title, newArticle.Content)
	}
	response.OkWithMessage("文章更新成功", c)
}
