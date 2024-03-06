package redis

import (
	"context"
	"gvb_server/global"
	"strconv"
)

const DiggPrefix = "digg"

// Digg 点赞某一篇文章
func Digg(id string) error {
	num, _ := global.Redis.HGet(context.Background(), DiggPrefix, id).Int()
	num++
	err := global.Redis.HSet(context.Background(), DiggPrefix, id, num).Err()
	return err
}

// GetDigg 获取某一篇文章的点赞数
func GetDigg(id string) int {
	num, _ := global.Redis.HGet(context.Background(), DiggPrefix, id).Int()
	return num
}

// GetDiggInfo 获取所有文章的点赞数
func GetDiggInfo() map[string]int {
	var DiggInfo = make(map[string]int)
	maps := global.Redis.HGetAll(context.Background(), DiggPrefix).Val()
	for k, v := range maps {
		num, _ := strconv.Atoi(v)
		DiggInfo[k] = num
	}
	return DiggInfo
}

// DiggClear 清空所有文章的点赞数
func DiggClear() error {
	err := global.Redis.Del(context.Background(), DiggPrefix).Err()
	return err
}
