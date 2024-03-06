package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time `json:"created_at"`           // 创建时间
	UpdatedAt time.Time `json:"updated_at"`           // 更新时间
}

type Page struct {
	PageNum int    `form:"page_num"` // 页码
	Key     string `form:"key"`      // 关键字
	Limit   int    `form:"limit"`    // 每页显示条数
	Sort    string `form:"sort"`     // 排序
}

type RemoveRequest struct {
	IDList []int `json:"id_list"`
}

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}
