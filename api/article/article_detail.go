package article

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/es"
	"gvb_server/service/redis"
)

func (this *ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}
	redis.Look(cr.ID)
	article, err := es.CommonDetail(cr.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(article, c)
}

func (this *ArticleApi) ArticleDetailByTitleView(c *gin.Context) {
	title := c.Query("title")
	article, err := es.CommonDetailByKeyword(title)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(article, c)
}
