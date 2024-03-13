package utils

import "gvb_server/global"

func PrintSystem(addr string) {
	global.Log.Infof("gvb_server 运行在：%s", "http://"+addr+"/api")
	global.Log.Infof("api文档 运行在：%s", "http://"+addr+"/swagger/index.html")
}
