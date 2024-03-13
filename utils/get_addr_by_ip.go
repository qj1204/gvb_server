package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"net"
)

func GetAddrByGin(c *gin.Context) (ip, addr string) {
	ip = c.ClientIP()
	addr = GetAddr(ip)
	return
}

func GetAddr(ip string) string {
	parseIP := net.ParseIP(ip)
	if IsInternalIP(parseIP) {
		return "内网IP"
	}
	record, err := global.AddrDB.City(parseIP)
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
