package redis

import (
	"context"
	"fmt"
	"gvb_server/global"
	"strings"
	"time"
)

const BindEmailPrefix = "bind_email_"

// BindEmail 绑定邮箱
func BindEmail(token, email, code string) error {
	err := global.Redis.Set(context.Background(), fmt.Sprintf("%s%s", BindEmailPrefix, token), email+"_"+code,
		time.Duration(global.Config.Redis.TTL)*time.Minute).Err()
	return err
}

// GetEmailAndCodeByToken 获取绑定邮箱
func GetEmailAndCodeByToken(token string) (string, string, error) {
	emailCode, err := global.Redis.Get(context.Background(), fmt.Sprintf("%s%s", BindEmailPrefix, token)).Result()
	if err != nil {
		return "", "", err
	}
	i := strings.LastIndex(emailCode, "_")
	return emailCode[:i], emailCode[i+1:], nil
}

// DelEmailAndCode 删除redis中的验证码
func DelEmailAndCode(token string) error {
	err := global.Redis.Del(context.Background(), fmt.Sprintf("%s%s", BindEmailPrefix, token)).Err()
	return err
}
