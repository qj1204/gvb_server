package user_service

import (
	"gvb_server/service/redis_service"
	"gvb_server/utils/jwts"
	"time"
)

func (UserService) Logout(claims *jwts.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redis_service.Logout(token, diff)
}
