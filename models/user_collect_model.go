package models

import "time"

// UserCollectModel 自定义用户收藏表，记录用户什么时候收藏了什么文章
type UserCollectModel struct {
	CreatedAt time.Time `json:"created_at"`
	UserID    uint      `gorm:"primaryKey"`
	UserModel UserModel `gorm:"foreignKey:UserID"`

	ArticleID string `gorm:"size:32"`
}
