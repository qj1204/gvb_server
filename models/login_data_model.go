package models

type LoginDataModel struct {
	MODEL

	// --------登录数据 belongs to 用户--------
	UserID    uint      `json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`

	IP       string `gorm:"size:20" json:"ip"`
	NickName string `gorm:"size:42" json:"nick_name"`
	Token    string `gorm:"size:256" json:"token"`
	Device   string `gorm:"size:256" json:"device"`
	Addr     string `gorm:"size:64" json:"addr"`
}
