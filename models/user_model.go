package models

import (
	"gvb_server/models/common/ctype"
)

// UserModel 用户表
type UserModel struct {
	MODEL
	NickName   string           `gorm:"size:36;comment:'昵称'" json:"nick_name"`               // 昵称
	UserName   string           `gorm:"size:36;comment:'用户名'" json:"user_name"`              // 用户名
	Password   string           `gorm:"size:128;comment:'密码'" json:"-"`                      // 密码
	Avatar     string           `gorm:"size:256;comment:'头像'" json:"avatar_id"`              // 头像id
	Email      string           `gorm:"size:128;comment:'邮箱'" json:"email"`                  // 邮箱
	Tel        string           `gorm:"size:18;comment:'电话'" json:"tel"`                     // 电话
	Addr       string           `gorm:"size:64;comment:'地址'" json:"addr"`                    // 地址
	Token      string           `gorm:"size:64;comment:'令牌'" json:"token"`                   // 令牌
	IP         string           `gorm:"size:20;comment:'IP'" json:"ip"`                      // IP
	Role       ctype.Role       `gorm:"size:4;default:1;comment:'角色'" json:"role"`           // 角色
	SignStatus ctype.SignStatus `gorm:"type=smallint(2)';comment:'登录方式'" json:"sign_status"` // 登录方式
}
