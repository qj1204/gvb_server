package jwt

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type JwtPayLoad struct {
	UserID uint `json:"user_id"` // 用户ID
	//Username string `json:"username"`  // 用户名
	NickName string `json:"nick_name"` // 昵称
	Role     int    `json:"role"`      // 角色 1：管理员 2：普通用户 3：游客
}

var MySecret []byte

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}
