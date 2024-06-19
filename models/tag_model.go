package models

// TagModel 标签表
type TagModel struct {
	MODEL
	Title string `gorm:"size:16;comment:'标签名称'" json:"title"` // 标签名称
}
