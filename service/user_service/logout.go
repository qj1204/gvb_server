package user_service

import (
	"gvb_server/service/redis_service"
	"gvb_server/utils/jwt"
	"time"
)

func (this *UserService) Logout(claims *jwt.CustomClaims, token string) error {
	// 需要计算距离现在的过期时间
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redis_service.Logout(token, diff)
}
