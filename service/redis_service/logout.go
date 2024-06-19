package redis_service

import (
	"context"
	"fmt"
	"gvb_server/global"
	"gvb_server/utils"
	"time"
)

const LogoutPrefix = "logout_"

// Logout 针对注销的操作
func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(context.Background(), fmt.Sprintf("%s%s", LogoutPrefix, token), "", diff).Err()
	return err
}

// CheckLogout 检查是否注销
func CheckLogout(token string) bool {
	keys := global.Redis.Keys(context.Background(), LogoutPrefix+"*").Val()
	if utils.InList(LogoutPrefix+token, keys) {
		return true
	}
	return false
}
