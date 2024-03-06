package redis

import (
	"context"
	"gvb_server/global"
	"strconv"
)

const LookPrefix = "look"

// Look 浏览某一篇文章
func Look(id string) error {
	num, _ := global.Redis.HGet(context.Background(), LookPrefix, id).Int()
	num++
	err := global.Redis.HSet(context.Background(), LookPrefix, id, num).Err()
	return err
}

// GetLook 获取某一篇文章的浏览量
func GetLook(id string) int {
	num, _ := global.Redis.HGet(context.Background(), LookPrefix, id).Int()
	return num
}

// GetLookInfo 获取所有文章的浏览量
func GetLookInfo() map[string]int {
	var LookInfo = make(map[string]int)
	maps := global.Redis.HGetAll(context.Background(), LookPrefix).Val()
	for k, v := range maps {
		num, _ := strconv.Atoi(v)
		LookInfo[k] = num
	}
	return LookInfo
}

// LookClear 清空redis中所有文章的浏览量
func LookClear() error {
	err := global.Redis.Del(context.Background(), LookPrefix).Err()
	return err
}
