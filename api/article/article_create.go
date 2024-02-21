package article

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/ctype"
	"gvb_server/models/common/response"
	"gvb_server/utils/jwt"
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
	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	userID := claims.UserID
	userNickName := claims.NickName
	// 处理content

	// 如果没有简介，就截取content的前100个字符
	if cr.Abstract == "" {
		// 汉字的截取不一样
		t := []rune(cr.Content)
		// 将content转为html， 并且过滤xss，以及获取中文内容
		cr.Abstract = string(t)
		if len(t) > 100 {
			cr.Abstract = string(t[:100])
		}
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
	err = global.DB.Model(models.UserModel{}).Where("id = ?", userID).Select("avatar").Scan(&avatar).Error
	if err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	article := models.ArticleModel{
		CreatedAt:    now,
		UpdatedAt:    now,
		Title:        cr.Title,
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		UserID:       userID,
		UserNickName: userNickName,
		UserAvatar:   avatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    bannerUrl,
		Tags:         cr.Tags,
	}

	err = article.Create()
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("文章发布成功", c)
}
