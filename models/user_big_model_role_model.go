package models

// UserBigModelRoleModel 用户选择的大模型角色表
type UserBigModelRoleModel struct {
	MODEL
	UserID    uint              `json:"userID"` // 用户id
	RoleID    uint              `json:"roleID"` // 角色id
	RoleModel BigModelRoleModel `gorm:"foreignKey:RoleID" json:"-"`
}
