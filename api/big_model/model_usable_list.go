package big_model

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/response"
)

// ModelUsableListView 可用的大模型列表接口
func (BigModelApi) ModelUsableListView(c *gin.Context) {
	response.OkWithData(global.Config.BigModel.ModelList, c)
	return
}
