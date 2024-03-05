package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
	"time"
)

type TagCreated struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	// 初始化数据库
	global.DB = core.InitGorm()

	tagTitleList := []string{"go", "java", "后端"}
	var tagCreatedList []TagCreated
	global.DB.Model(&models.TagModel{}).Select("title", "created_at").Where("title in ?", tagTitleList).Scan(&tagCreatedList)
	tagMap := make(map[string]time.Time)
	for _, tag := range tagCreatedList {
		tagMap[tag.Title] = tag.CreatedAt
	}
	fmt.Println(tagMap)
}
