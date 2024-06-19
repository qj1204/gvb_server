package big_model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
	"regexp"
)

type AutoReplyUpdateRequest struct {
	ID           uint   `json:"id"`
	Name         string `json:"name" binding:"required"`               // 规则名称
	Mode         int    `json:"mode" binding:"required,oneof=1 2 3 4"` // 匹配模式 1 精确匹配，2 模糊匹配，3 前缀匹配，4 正则匹配
	Rule         string `json:"rule" binding:"required"`               // 匹配规则
	ReplyContent string `json:"replyContent" binding:"required"`       // 回复内容
}

// AutoReplyUpdateView 增加和修改
func (BigModelApi) AutoReplyUpdateView(c *gin.Context) {
	var cr AutoReplyUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithValidError(err, c)
		return
	}

	// 校验正则是否写错
	if cr.Mode == 4 {
		_, err := regexp.Compile(cr.Rule)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("正则表达式错误 %s", err.Error()), c)
			return
		}
	}

	if cr.ID == 0 {
		// 增加
		var arm models.AutoReplyModel
		err = global.DB.Take(&arm, "name = ?", cr.Name).Error
		if err == nil {
			response.FailWithMessage("规则名称不能相同", c)
			return
		}
		err = global.DB.Create(&models.AutoReplyModel{
			Name:         cr.Name,
			Mode:         cr.Mode,
			Rule:         cr.Rule,
			ReplyContent: cr.ReplyContent,
		}).Error
		if err != nil {
			logrus.Errorf("数据添加失败 err：%s, 原始数据内容 %#v", err, cr)
			response.FailWithMessage("数据添加失败", c)
			return
		}
		response.OkWithMessage("自动回复添加成功", c)
		return
	}
	var arm models.AutoReplyModel
	err = global.DB.Take(&arm, cr.ID).Error
	if err != nil {
		response.FailWithMessage("记录不存在", c)
		return
	}
	// name不能和已有的重复
	var arm1 models.AutoReplyModel
	err = global.DB.Take(&arm1, "name = ? and id <> ?", cr.Name, cr.ID).Error
	if err == nil {
		response.FailWithMessage("和已有的规则名称重复", c)
		return
	}
	err = global.DB.Model(&arm).Updates(map[string]any{
		"name":          cr.Name,
		"mode":          cr.Mode,
		"rule":          cr.Rule,
		"reply_content": cr.ReplyContent,
	}).Error
	if err != nil {
		logrus.Errorf("数据更新失败 err：%s, 原始数据内容 %#v", err, cr)
		response.FailWithMessage("数据更新失败", c)
		return
	}
	response.OkWithMessage("自动回复更新成功", c)
	return
}
