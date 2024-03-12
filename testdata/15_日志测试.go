package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/plugins/log_stash"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.DB = core.InitGorm()
	log := log_stash.NewLog("192.168.100.10", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuaWNrX25hbWUiOiLnrqHnkIblkZgiLCJyb2xlIjoxLCJleHAiOjE3MTAzODY2MzkuMTc4OTI0LCJpc3MiOiJxaWFuamluIn0.gALzs6izuUZhmd0AKYdy71ezRmSauwBsj0he9U4ZnZk")
	// fmt.Printf("%p\n", log)
	log.Debug("哈哈哈")
}
