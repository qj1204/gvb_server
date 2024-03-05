package redis

import (
	"context"
	"fmt"
	"gvb_server/global"
	"gvb_server/utils"
	"strings"
	"time"
)

const (
	logoutPrefix    = "logout_"
	bindEmailPrefix = "bind_email_"
)

// Logout 针对注销的操作
func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(context.Background(), fmt.Sprintf("%s%s", logoutPrefix, token), "", diff).Err()
	return err
}

// CheckLogout 检查是否注销
func CheckLogout(token string) bool {
	keys := global.Redis.Keys(context.Background(), logoutPrefix+"*").Val()
	if utils.InList(logoutPrefix+token, keys) {
		return true
	}
	return false
}

// BindEmail 绑定邮箱
func BindEmail(token, email, code string) error {
	err := global.Redis.Set(context.Background(), fmt.Sprintf("%s%s", bindEmailPrefix, token), email+"_"+code,
		time.Duration(global.Config.Redis.TTL)*time.Minute).Err()
	return err
}

// GetEmailAndCodeByToken 获取绑定邮箱
func GetEmailAndCodeByToken(token string) (string, string, error) {
	emailCode, err := global.Redis.Get(context.Background(), fmt.Sprintf("%s%s", bindEmailPrefix, token)).Result()
	if err != nil {
		return "", "", err
	}
	i := strings.LastIndex(emailCode, "_")
	return emailCode[:i], emailCode[i+1:], nil
}

// DelEmailAndCode 删除redis中的验证码
func DelEmailAndCode(token string) error {
	err := global.Redis.Del(context.Background(), fmt.Sprintf("%s%s", bindEmailPrefix, token)).Err()
	return err
}
