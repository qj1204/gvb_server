package log_stash

import "time"

// LogStashModel 日志
type LogStashModel struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	IP        string    `gorm:"size:32" json:"ip"`
	Addr      string    `gorm:"size:64" json:"addr"`
	Level     Level     `gorm:"size:4" json:"level"`     // 日志级别
	Content   string    `gorm:"size:128" json:"content"` // 日志消息内容
	UserID    uint      `json:"user_id"`                 // 用户ID，不设外键，需要在查询的时候手动关联
}
