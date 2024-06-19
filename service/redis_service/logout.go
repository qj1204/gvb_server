package redis_service

import (
	"context"
	"gvb_server/global"
	"gvb_server/utils"
	"time"
)

const prefix = "logout_"

// Logout 针对注销的操作
func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(context.Background(), prefix+token, "", diff).Err()
	return err
}

func CheckLogout(token string) bool {
	keys := global.Redis.Keys(context.Background(), prefix+"*").Val()
	if utils.InList(prefix+token, keys) {
		return true
	}
	return false
}
