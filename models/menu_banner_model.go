package models

// MenuBannerModel 菜单轮播图关联表
type MenuBannerModel struct {
	// --------菜单 自定义多对多 Banner--------
	MenuID   uint `gorm:"primaryKey;comment:菜单id" json:"menu_id"`
	BannerID uint `gorm:"primaryKey;comment:banner图id" json:"banner_id"`

	MenuModel   MenuModel   `gorm:"foreignKey:MenuID"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID"`

	Sort int `gorm:"size:10;comment:序号" json:"sort"`
}
