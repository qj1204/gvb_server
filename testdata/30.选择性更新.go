package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"reflect"
)

type RoleUpdateRequest struct {
	ID        uint    `json:"id"`
	Name      *string `json:"name" structs:"name"`            // 角色名称
	Enable    *bool   `json:"enable" structs:"enable"`        // 是否启用
	Icon      *string `json:"icon" structs:"icon"`            // 可以选择系统默认的一些，也可以图片上传
	Abstract  *string `json:"abstract" structs:"abstract"`    // 简介
	Scope     *int    `json:"scope" structs:"scope"`          // 消耗的积分
	Prologue  *string `json:"prologue" structs:"prologue"`    // 开场白
	Prompt    *string `json:"prompt" structs:"prompt"`        // 设定词
	AutoReply *bool   `json:"autoReply" structs:"auto_reply"` // 自动回复
	TagList   *[]uint `json:"tagList" structs:"tagList"`      // 标签的id列表
}

func main() {
	var cr RoleUpdateRequest
	data := `
	{
		"id": 2,
		"name": "IT专家",
		"enable": true,
		"icon": "/uploads/xxx.png",
		"abstract": "一个帮助你解决it问题的小帮手",
		"scope": 3
	}
	`
	json.Unmarshal([]byte(data), &cr)

	maps := structs.Map(cr)
	var mps = map[string]any{}
	for s, i := range maps {
		if s == "ID" {
			continue
		}
		if i == nil {
			continue
		}
		val := reflect.ValueOf(i)
		if val.Kind() == reflect.Ptr && val.IsNil() {
			continue
		}
		mps[s] = val.Elem().Interface()
	}

	for s, a := range mps {
		fmt.Println(s, a)
	}
}
