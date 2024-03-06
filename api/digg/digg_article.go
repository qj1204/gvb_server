package digg

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/redis"
)

func (this *DiggApi) DiggArticleView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}
	// 这里可以校验文章id
	// 查es
	redis.Digg(cr.ID)
	response.OkWithMessage("文章点赞成功", c)
}