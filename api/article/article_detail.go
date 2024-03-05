package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/models/common/response"
	"gvb_server/service/es"
)

func (this *ArticleApi) ArticleDetailView(c *gin.Context) {
	id := c.Param("id")
	article, err := es.CommonDetail(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(article, c)
}

func (this *ArticleApi) ArticleDetailByTitleView(c *gin.Context) {
	title := c.Query("title")
	fmt.Println(title)
	article, err := es.CommonDetailByKeyword(title)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(article, c)
}
