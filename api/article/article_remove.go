package article

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
)

type IDListRequest struct {
	IDList []string `json:"id_list"`
}

func (this *ArticleApi) ArticleRemoveView(c *gin.Context) {
	var cr IDListRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	bulkService := global.ESClient.Bulk().Index(models.ArticleModel{}.Index()).Refresh("true")
	for _, id := range cr.IDList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
	}
	res, err := bulkService.Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("删除文章失败", c)
		return
	}
	response.OkWithMessage(fmt.Sprintf("共删除%d篇文章", len(res.Succeeded())), c)
}
