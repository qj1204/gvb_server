package core

import (
	"github.com/oschwald/geoip2-golang"
	"gvb_server/global"
)

func InitAddrDB() {
	db, err := geoip2.Open("uploads/system/GeoLite2-City.mmdb")
	if err != nil {
		global.Log.Fatal("ip地址数据库加载失败", err)
	}
	global.AddrDB = db
}
