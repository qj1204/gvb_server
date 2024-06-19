package models

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"os"
)

// BannerModel 图片表
type BannerModel struct {
	MODEL
	Path        string            `gorm:"comment:'图片路径'" json:"path"`                 // 图片路径
	Hash        string            `gorm:"comment:'图片哈希'" json:"hash"`                 // 图片哈希，用于判断图片是否重复
	Name        string            `gorm:"size:38;comment:'图片名称'" json:"name"`         // 图片名称
	ImageType   ctype.ImageType   `gorm:"default:1;comment:'图片类型'" json:"image_type"` // 图片类型(本地/七牛云)
	MenusBanner []MenuBannerModel `gorm:"foreignKey:BannerID" json:"-"`
}

// BeforeDelete 删库前，需要删除本地图片
func (this *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	// 本地图片，删除，还要删除本地的存储
	if this.ImageType == ctype.Local {
		err = os.Remove(this.Path)
		if err != nil {
			global.Log.Error(err)
			return
		}
	}
	return
}
