package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"gvb_server/global"
)

// ParseToken 解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Log.Error(fmt.Sprintf("解析token失败：%s", err.Error()))
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效的token")
}
