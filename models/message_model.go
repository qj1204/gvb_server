package models

// MessageModel 消息表
type MessageModel struct {
	MODEL

	// --------消息 belongs to 发送者--------
	SendUserID    uint      `gorm:"primaryKey" json:"send_user_id"` // 发送者ID
	SendUserModel UserModel `gorm:"foreignKey:SendUserID" json:"-"`

	SendUserNickName string `gorm:"size:36" json:"send_user_nick_name"` // 发送者昵称
	SendUserAvatar   string `json:"send_user_avatar"`                   // 发送者头像

	// --------消息 belongs to 接收者--------
	ReceiveUserID    uint      `gorm:"primaryKey" json:"receive_user_id"` // 接收者ID
	ReceiveUserModel UserModel `gorm:"foreignKey:ReceiveUserID" json:"-"`

	ReceiveUserNickName string `gorm:"size:36" json:"receive_user_nick_name"` // 接收者昵称
	ReceiveUserAvatar   string `json:"receive_user_avatar"`                   // 接收者头像
	IsRead              bool   `gorm:"default:false" json:"is_read"`          // 是否已读
	Content             string `json:"content"`                               // 消息内容
}
