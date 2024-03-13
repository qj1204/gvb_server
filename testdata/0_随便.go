package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"slices"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.DB = core.InitGorm()
	global.Redis = core.ConnectRedis()
	oldTags := []string{"java", "go"}
	oldTags = slices.DeleteFunc(oldTags, func(s string) bool {
		return s == "java" || s == "go"
	})
	fmt.Println(oldTags)
}
