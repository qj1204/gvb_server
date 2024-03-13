package main

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"net"
)

func main() {
	fmt.Println(GetAddr("39.99.248.35"))
	fmt.Println(GetAddr("1113.200.174.21"))
	fmt.Println(GetAddr("192.168.100.10"))
	fmt.Println(GetAddr("172.10.56.5"))
}

func GetAddr(ip string) string {
	parseIP := net.ParseIP(ip)
	if IsInternalIP(parseIP) {
		return "内网IP"
	}
	db, _ := geoip2.Open("static/system/GeoLite2-City.mmdb")
	defer db.Close()
	record, err := db.City(parseIP)
	if err != nil {
		return "错误的IP地址"
	}
	var provice string
	if len(record.Subdivisions) > 0 { // 有些IP地址没有省份信息
		provice = record.Subdivisions[0].Names["zh-CN"]
	}
	city := record.City.Names["zh-CN"]
	return fmt.Sprintf("%s-%s", provice, city)
}

func IsInternalIP(ip net.IP) bool {
	return ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() || ip.IsPrivate()
}
