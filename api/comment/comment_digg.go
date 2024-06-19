package comment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/service/redis_service"
)

type CommentIDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

// CommentDiggView 评论点赞
// @Tags 评论管理
// @Summary 评论点赞
// @Description 评论点赞
// @Param id path int  true  "id"
// @Router /api/comments/digg/{id} [get]
// @Produce json
// @Success 200 {object} response.Response{}
func (CommentApi) CommentDiggView(c *gin.Context) {
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
	redis_service.NewCommentDiggCount().Set(fmt.Sprintf("%d", cr.ID))
	response.OkWithMessage("评论点赞成功", c)
}
