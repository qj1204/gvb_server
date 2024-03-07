package article

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/es"
	"gvb_server/utils/jwt"
	"time"
)

// ArticleCollectCreateView 收藏文章或者取消收藏
func (this *ArticleApi) ArticleCollectCreateView(c *gin.Context) {
	var cr models.ESIDRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	article, err := es.CommonDetail(cr.ID)
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("文章不存在", c)
		return
	}

	var num = 1
	var collect models.UserCollectModel
	err = global.DB.Take(&collect, "user_id = ? and article_id = ?", claims.UserID, cr.ID).Error
	if err == nil { // 取消收藏
		global.DB.Delete(&collect)
		// 给文章的收藏数-1
		num = -1
	} else { // 收藏文章
		global.DB.Create(&models.UserCollectModel{
			CreatedAt: time.Now(),
			UserID:    claims.UserID,
			ArticleID: cr.ID,
		})
	}
	// 更新es中文章的收藏数
	err = es.ArticleUpdate(cr.ID, map[string]any{"collects_count": article.CollectsCount + num})
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("收藏数更新失败", c)
		return
	}
	if num == 1 {
		response.OkWithMessage("收藏文章成功", c)
		return
	}
	response.OkWithMessage("取消收藏成功", c)
}
