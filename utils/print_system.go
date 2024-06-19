package utils

import "gvb_server/global"

func PrintSystem() {
	ip := global.Config.System.Host
	port := global.Config.System.Port
	global.Log.Infof("gvb_server 运行在：http://%s:%d/api", ip, port)
	global.Log.Infof("api文档 运行在：http://%s:%d/swagger/index.html#", ip, port)
}
