package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
)

func main() {
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()

	reply := models.AutoReplyModel{}
	model := reply.ValidAutoReply("请问你是谁？")
	fmt.Println(model)
}
