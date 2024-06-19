package ctype

import "encoding/json"

// 参考gorm代码的10_枚举.go

type Role int

const (
	PermissionAdmin Role = 1 // 管理员
	PermissionUser  Role = 2 // 普通用户
	PermissionGuest Role = 3 // 游客
	PermissionBan   Role = 4 // 封禁用户
)

func (this Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.String())
}

func (this Role) String() string {
	var s string
	switch this {
	case PermissionAdmin:
		s = "管理员"
	case PermissionUser:
		s = "普通用户"

	case PermissionGuest:
		s = "游客"
	case PermissionBan:
		s = "封禁用户"
	default:
		s = "未知"
	}
	return s
}
