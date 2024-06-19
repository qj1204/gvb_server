package models

// MessageModel 消息表
type MessageModel struct {
	MODEL

	// --------消息 belongs to 发送者--------
	SendUserID    uint      `gorm:"primaryKey;comment:发送人id" json:"send_user_id"` // 发送人ID
	SendUserModel UserModel `gorm:"foreignKey:SendUserID" json:"-"`

	SendUserNickName string `gorm:"size:36;comment:发送人昵称" json:"send_user_nick_name"` // 发送人昵称
	SendUserAvatar   string `gorm:"comment:发送人头像" json:"send_user_avatar"`            // 发送人头像

	// --------消息 belongs to 接收者--------
	ReceiveUserID    uint      `gorm:"primaryKey;;comment:接收人id" json:"receive_user_id"` // 接收人ID
	ReceiveUserModel UserModel `gorm:"foreignKey:ReceiveUserID" json:"-"`

	ReceiveUserNickName string `gorm:"size:36;comment:接收人昵称" json:"receive_user_nick_name"` // 接收人昵称
	ReceiveUserAvatar   string `gorm:"comment:接收人头像"  json:"receive_user_avatar"`           // 接收人头像
	IsRead              bool   `gorm:"default:false;comment:接收人是否查看" json:"is_read"`        // 是否已读
	Content             string `gorm:"comment:消息内容" json:"content"`                         // 消息内容
}
