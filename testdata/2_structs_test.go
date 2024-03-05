package main

import (
	"fmt"
	"github.com/fatih/structs"
	"testing"
)

type AdvertRequest struct {
	Title  string `gorm:"size:32" json:"title" binding:"required" msg:"请输入广告标题" structs:"title"` // 广告标题
	Href   string `json:"href" binding:"required,url" msg:"广告链接非法" structs:"href"`               // 广告链接
	Image  string `json:"image" binding:"required,url" msg:"广告图片地址" structs:"image"`             // 广告图片
	IsShow *bool  `json:"is_show" binding:"required" msg:"请选择是否展示" structs:"is_show"`            // 是否显示
}

func TestStructs(t *testing.T) {
	b := true
	a1 := AdvertRequest{
		Title:  "test",
		Href:   "http://www.baidu.com",
		Image:  "http://www.baidu.com",
		IsShow: &b,
	}
	m := structs.Map(a1)
	fmt.Println(m)
}
