package ctype

import "encoding/json"

// 参考gorm代码的10_枚举.go

type SignStatus int

const (
	QQStatus  SignStatus = 1 // QQ登录
	SignGitee SignStatus = 2 // Gitee登录
	SignEmail SignStatus = 3 // 邮箱登录
)

func (this SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.String())
}

func (this SignStatus) String() string {
	var s string
	switch this {
	case QQStatus:
		s = "QQ登录"
	case SignGitee:
		s = "Gitee登录"
	case SignEmail:
		s = "邮箱登录"
	default:
		s = "未知"
	}
	return s
}
