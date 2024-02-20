package models

// MenuBannerModel 菜单轮播图关联表
type MenuBannerModel struct {
	// --------menu 自定义多对多 Banner--------
	MenuID   uint `gorm:"primaryKey" json:"menu_id"`
	BannerID uint `gorm:"primaryKey" json:"banner_id"`

	MenuModel   MenuModel   `gorm:"foreignKey:MenuID"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID"`

	Sort int `gorm:"size:10" json:"sort"`
}
