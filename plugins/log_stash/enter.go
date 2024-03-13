package log_stash

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/utils"
	"gvb_server/utils/jwt"
)

type Log struct {
	ip     string `json:"ip"`
	addr   string `json:"addr"`
	userID uint   `json:"user_id"`
}

func NewLog(ip, token string) *Log {
	// 解析token，拿到用户id
	claims, err := jwt.ParseToken(token)
	var userID uint
	if err == nil {
		userID = claims.UserID
	}
	return &Log{ip: ip, addr: utils.GetAddr(ip), userID: userID}
}

func NewLogByGin(c *gin.Context) *Log {
	token := c.Request.Header.Get("token")
	return NewLog(c.ClientIP(), token)
}

func (this *Log) Debug(content string) {
	this.Send(DebugLevel, content)
}

func (this *Log) Info(content string) {
	this.Send(InfoLevel, content)
}

func (this *Log) Warn(content string) {
	this.Send(WarnLevel, content)
}

func (this *Log) Error(content string) {
	this.Send(ErrorLevel, content)
}

// Send 日志消息入库
func (this *Log) Send(level Level, content string) {
	if err := global.DB.Create(&LogStashModel{
		IP:      this.ip,
		Addr:    this.addr,
		Level:   level,
		Content: content,
		UserID:  this.userID,
	}).Error; err != nil {
		global.Log.Error(err)
	}
	fmt.Println(this.ip, this.addr, this.userID, level, content)
}

//func Debug(content string) {
//	std.Debug(content)
//}
//
//func Info(content string) {
//	std.Info(content)
//}
//
//func Warn(content string) {
//	std.Warn(content)
//}
//
//func Error(content string) {
//	std.Error(content)
//}
