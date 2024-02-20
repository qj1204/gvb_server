package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time `json:"created_at"`           // 创建时间
	UpdatedAt time.Time `json:"updated_at"`           // 更新时间
}

type Page struct {
	PageNum int    `form:"page_num"`
	Key     string `form:"key"`
	Limit   int    `form:"limit"`
	Sort    string `form:"sort"`
}

type RemoveRequest struct {
	IDList []int `json:"id_list"`
}
