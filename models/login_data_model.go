package models

import "gvb_server/models/ctype"

type LoginDataModel struct {
	MODEL

	// --------登录数据 belongs to 用户--------
	UserID    uint      `gorm:"comment:用户id" json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`

	IP        string           `gorm:"size:20;comment:ip" json:"ip"`
	NickName  string           `gorm:"size:42;comment:昵称" json:"nick_name"`
	Token     string           `gorm:"size:256;comment:token" json:"token"`
	Device    string           `gorm:"size:256;comment:登录设备" json:"device"`
	Addr      string           `gorm:"size:64;comment:地址" json:"addr"`
	LoginType ctype.SignStatus `gorm:"size:smallint(6);comment:登录方式，1QQ，3邮箱" json:"login_type"` // 登录方式
}
