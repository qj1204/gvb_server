package comment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/redis"
)

type CommentIDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

func (this *CommentApi) CommentDiggView(c *gin.Context) {
	var cr CommentIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	var comment models.CommentModel
	err = global.DB.Take(&comment, cr.ID).Error
	if err != nil {
		response.FailWithMessage("评论不存在", c)
		return
	}
	// TODO: 评论点赞同步到数据库
	redis.NewCommentDiggCount().Set(fmt.Sprintf("%d", cr.ID))
	response.OkWithMessage("评论点赞成功", c)
}
