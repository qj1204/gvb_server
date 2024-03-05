package ctype

import "encoding/json"

// 参考gorm代码的10_枚举.go

type ImageType int

const (
	Local ImageType = 1 // 本地
	QiNiu ImageType = 2 // 七牛云
)

func (this ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.String())
}

func (this ImageType) String() string {
	var s string
	switch this {
	case Local:
		s = "本地"
	case QiNiu:
		s = "七牛云"
	default:
		s = "未知"
	}
	return s
}
