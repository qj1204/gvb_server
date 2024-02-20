package models

import "gvb_server/models/common/ctype"

// MenuModel 菜单表
type MenuModel struct {
	MODEL
	Title        string      `gorm:"size:32;comment:'菜单标题'" json:"title"`      // 菜单标题
	Path         string      `gorm:"size:32;comment:'菜单路径'" json:"path"`       // 路径
	Slogan       string      `gorm:"size:64;comment:'标语'" json:"slogan"`       // 标语
	Abstract     ctype.Array `gorm:"type:string;comment:'简介'" json:"abstract"` // 简介
	AbstractTime int         `gorm:"comment:'简介的切换时间'" json:"abstract_time"`   // 简介的切换时间

	// --------菜单 自定义多对多 Banner--------
	Banners []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;joinReferences:BannerID" json:"banners"` // 菜单图片

	BannerTime int `gorm:"comment:'菜单的切换时间'" json:"banner_time"` // 菜单的切换时间，为0表示不切换
	Sort       int `gorm:"size:10;comment:'菜单的顺序'" json:"sort"`  // 菜单的顺序
}
