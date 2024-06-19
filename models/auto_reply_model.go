package models

import (
	"gvb_server/global"
	"regexp"
	"strings"
)

// AutoReplyModel 自动回复
type AutoReplyModel struct {
	MODEL
	Name         string `gorm:"size:32" json:"name"`           // 规则名称
	Mode         int    `json:"mode"`                          // 匹配模式 1 精确匹配，2 模糊匹配，3 前缀匹配，4 正则匹配
	Rule         string `gorm:"size:64" json:"rule"`           // 匹配规则
	ReplyContent string `gorm:"size:1024" json:"replyContent"` // 回复内容
}

// ValidAutoReply 是否命中自动回复
func (AutoReplyModel) ValidAutoReply(content string) *AutoReplyModel {
	var list []AutoReplyModel
	global.DB.Find(&list)
	for _, model := range list {
		switch model.Mode {
		case 1:
			// 精确
			if model.Rule == content {
				return &model
			}
		case 2:
			// 包含
			if strings.Contains(content, model.Rule) {
				return &model
			}
		case 3:
			// 前缀
			if strings.HasPrefix(content, model.Rule) {
				return &model
			}
		case 4:
			// 正则
			regex, _ := regexp.Compile(model.Rule)
			if regex.MatchString(content) {
				return &model
			}
		}
	}
	return nil
}
