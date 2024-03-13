package main

import (
	"gvb_server/core"
	_ "gvb_server/docs"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
	"gvb_server/service/cron"
	"gvb_server/utils"
)

// @title gvb_server API文档
// @version 1.0
// @description gvb_server API文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 初始化数据库
	global.DB = core.InitGorm()

	core.InitAddrDB()
	defer global.AddrDB.Close()

	// 命令行参数绑定（在连接数据库之后，链接路由之前）
	option := flag.Parse()
	if flag.IsWebStop(option) {
		// 如果是停止web服务，则执行相应的操作
		flag.SwitchOption(option)
		return
	}

	// 连接redis
	global.Redis = core.ConnectRedis()
	// 连接es
	global.ESClient = core.EsConnect()
	// 初始化定时任务
	cron.CronInit()

	addr := global.Config.System.GetAddr()
	utils.PrintSystem(addr)

	err := routers.InitRouter().Run(addr)
	if err != nil {
		global.Log.Fatal(err.Error())
	}
}
