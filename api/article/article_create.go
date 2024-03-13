package article

import "C"
import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/ctype"
	"gvb_server/models/common/response"
	"gvb_server/service/es"
	"gvb_server/utils/jwt"
	"gvb_server/utils/random"
	"strings"
	"time"
)

type ArticleRequest struct {
	Title    string      `json:"title" binding:"required" msg:"文章标题必填"`   // 文章标题
	Abstract string      `json:"abstract"`                                // 文章简介
	Content  string      `json:"content" binding:"required" msg:"文章内容必填"` // 文章内容
	Category string      `json:"category"`                                // 文章分类
	Source   string      `json:"source"`                                  // 文章来源
	Link     string      `json:"link"`                                    // 原文链接
	BannerID uint        `json:"banner_id"`                               // 文章封面ID
	Tags     ctype.Array `json:"tags"`                                    // 文章标签
}

func (this *ArticleApi) ArticleCreateView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	// 判断标题是否重复
	_, err = es.CommonDetailByKeyword(cr.Title)
	if err == nil {
		response.FailWithMessage("文章已存在", c)
		return
	}

	// 处理content（将content转为html，并且过滤xss攻击，以及获取中文内容）
	unsafe := blackfriday.MarkdownCommon([]byte(cr.Content))
	// 是不是有script标签
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	nodes := doc.Find("script").Nodes
	if len(nodes) > 0 {
		// 过滤xss攻击
		doc.Find("script").Remove()
		// html转markdown
		converter := md.NewConverter("", true, nil)
		html, _ := doc.Html()
		markdown, _ := converter.ConvertString(html)
		cr.Content = markdown
	}

	// 如果没有简介，就截取content的前100个字符
	if cr.Abstract == "" {
		// 汉字的截取不一样
		t := []rune(doc.Text())
		cr.Abstract = string(t)
		if len(t) > 100 {
			cr.Abstract = string(t[:100])
		}
	}

	// 没传banner_id的话，就随机选一张
	if cr.BannerID == 0 {
		var bannerIDList []uint
		global.DB.Model(models.BannerModel{}).Select("id").Scan(&bannerIDList)
		if len(bannerIDList) == 0 {
			response.FailWithMessage("数据库中没有banner数据", c)
			return
		}
		cr.BannerID = random.ListID(bannerIDList)
	}

	// 查banner_id对应的banner_url
	var bannerUrl string
	err = global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl).Error
	if err != nil {
		response.FailWithMessage("banner不存在", c)
		return
	}

	// 查用户头像
	var avatar string
	err = global.DB.Model(models.UserModel{}).Where("id = ?", claims.UserID).Select("avatar").Scan(&avatar).Error
	if err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	article := models.ArticleModel{
		CreatedAt:    now,
		UpdatedAt:    now,
		Title:        cr.Title,
		Keyword:      cr.Title,
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		UserID:       claims.UserID,
		UserNickName: claims.NickName,
		UserAvatar:   avatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    bannerUrl,
		Tags:         cr.Tags,
	}

	err = article.InsertArticle()
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 同步文章数据到全文搜索索引
	go es.AsyncArticleByFullText(article.ID, article.Title, article.Content)
	response.OkWithMessage("文章发布成功", c)
}
