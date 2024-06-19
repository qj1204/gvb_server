package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

// ModelUsableListView 可用的大模型列表接口
// @Tags 大模型管理
// @Summary 可用的大模型列表接口
// @Description 可用的大模型列表接口
// @Router /api/big_model/usable [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]config.ModelOption}
func (BigModelApi) ModelUsableListView(c *gin.Context) {
	res.OkWithData(global.Config.BigModel.ModelList, c)
}
