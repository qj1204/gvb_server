package models

// UserScopeModel 用户积分表
type UserScopeModel struct {
	MODEL
	UserID uint `json:"userID"`
	Scope  int  `json:"scope"`
	Status bool `json:"status"`
}
