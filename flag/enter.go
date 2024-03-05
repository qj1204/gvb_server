package flag

import (
	sys_flag "flag"
	"github.com/fatih/structs"
	"gvb_server/core"
	"gvb_server/global"
)

type Option struct {
	DB   bool
	User string // -u admin  -u user
	ES   string // -es create  -es delete
}

// Parse 解析命令行参数
func Parse() Option {
	// 如果有-db，返回值为true，否则为false
	db := sys_flag.Bool("db", false, "初始化数据库表")
	// 如果有-u，返回值为-u后面的值，否则为""
	user := sys_flag.String("u", "", "创建用户")
	// 如果有-es，返回值为-es后面的值，否则为""
	es := sys_flag.String("es", "", "es")
	// 解析命令行参数
	sys_flag.Parse()
	return Option{ // 因为上面的返回值都是指针，所以这里要加*
		DB:   *db,
		User: *user,
		ES:   *es,
	}
}

// IsWebStop 是否停止web服务
func IsWebStop(option Option) (f bool) {
	maps := structs.Map(option)
	for _, v := range maps {
		// 只要有一个值不为空，就返回true
		switch val := v.(type) {
		case string:
			if val != "" {
				f = true
			}
		case bool:
			if val != false {
				f = true
			}
		}
	}
	return
}

// SwitchOption 根据命令行参数执行相应的操作
func SwitchOption(option Option) {
	if option.DB {
		// 有-db参数，初始化数据库表
		MakeMigrations()
		return
	}
	if option.User == "admin" || option.User == "user" {
		// 有-u参数，创建用户
		CreateUser(option.User)
		return
	}
	if option.ES == "create" {
		// 有-es参数，连接es
		global.ESClient = core.EsConnect()
		EsCreateIndex()
		return
	}
	// 如果没有参数，打印帮助信息
	sys_flag.Usage()
}
