package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id,select($any)"` // 主键ID
	CreatedAt time.Time `json:"created_at,select($any)"`           // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                        // 更新时间
}

type Page struct {
	PageNum int    `form:"page_num"` // 页码
	Key     string `form:"key"`      // 关键字
	Limit   int    `form:"limit"`    // 每页显示条数
	Sort    string `form:"sort"`     // 排序
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}

type IDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

type ESIDListRequest struct {
	IDList []string `json:"id_list" binding:"required"`
}

type Options[T any] struct {
	Label string `json:"label"`
	Value T      `json:"value"`
}
