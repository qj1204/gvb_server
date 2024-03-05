package jwt

import (
	"github.com/dgrijalva/jwt-go/v4"
	"gvb_server/global"
	"time"
)

// GenerateToken 生成token
func GenerateToken(user JwtPayLoad) (string, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), // 默认过期时间
			Issuer:    global.Config.Jwt.Issuer,                                                     // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(MySecret)
}
