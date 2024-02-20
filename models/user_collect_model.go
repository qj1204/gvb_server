package models

import "time"

// UserCollectModel 自定义用户收藏表，记录用户什么时候收藏了什么文章
type UserCollectModel struct {
	// --------用户 自定义多对多 收藏文章--------
	UserID    uint `gorm:"primaryKey"`
	ArticleID uint `gorm:"primaryKey"`

	UserModel UserModel `gorm:"foreignKey:UserID"`

	CreateAt time.Time `json:"create_at"`
}
