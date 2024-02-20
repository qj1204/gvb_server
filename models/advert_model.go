package models

// AdvertModel 广告表
type AdvertModel struct {
	MODEL
	Title  string `gorm:"size:32" json:"title"` // 广告标题
	Href   string `json:"href"`                 // 广告链接
	Image  string `json:"image"`                // 广告图片
	IsShow bool   `json:"is_show"`              // 是否显示
}
