package main

import (
	"fmt"
	"github.com/cc14514/go-geoip2"
	geoip2db "github.com/cc14514/go-geoip2-db"
	"net"
	"time"
)

var db *geoip2.DBReader

func init() {
	db, _ = geoip2db.NewGeoipDbByStatik()
	//defer db.Close()
}
func main() {
	now := time.Now()
	fmt.Println(GetAddr("113.246.158.106"))

	fmt.Println(GetAddr("1131.246.158.106"))
	fmt.Println(GetAddr("123.34.5.106"))
	fmt.Println(GetAddr("192.168.100.1"))
	fmt.Println(GetAddr("172.10.56.5"))
	fmt.Println(GetAddr("127.0.0.1"))
	fmt.Println(time.Since(now))
	defer db.Close()
}

func GetAddr(ip string) string {
	parseIP := net.ParseIP(ip)
	if IsIntranetIP(parseIP) {
		return "内网地址"
	}
	record, err := db.City(net.ParseIP(ip))
	if err != nil {
		return "错误的地址"
	}
	var province string
	if len(record.Subdivisions) > 0 {
		province = record.Subdivisions[0].Names["zh-CN"]
	}
	city := record.City.Names["zh-CN"]
	return fmt.Sprintf("%s-%s", province, city)
}

func IsIntranetIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return true
	}
	// 192.168
	// 172.16 - 172.31
	// 10
	// 169.254
	return (ip4[0] == 192 && ip4[1] == 168) ||
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 32) ||
		(ip4[0] == 10) ||
		(ip4[0] == 169 && ip4[1] == 254)
}
