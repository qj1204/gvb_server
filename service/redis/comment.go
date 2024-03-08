package redis

import (
	"context"
	"gvb_server/global"
	"strconv"
)

const CommentPrefix = "comment"

// Comment 评论某一篇文章
func Comment(id string) error {
	num, _ := global.Redis.HGet(context.Background(), CommentPrefix, id).Int()
	num++
	err := global.Redis.HSet(context.Background(), CommentPrefix, id, num).Err()
	return err
}

// GetComment 获取某一篇文章的评论数
func GetComment(id string) int {
	num, _ := global.Redis.HGet(context.Background(), CommentPrefix, id).Int()
	return num
}

// GetCommentInfo 获取所有文章的评论数
func GetCommentInfo() map[string]int {
	var DiggInfo = make(map[string]int)
	maps := global.Redis.HGetAll(context.Background(), CommentPrefix).Val()
	for k, v := range maps {
		num, _ := strconv.Atoi(v)
		DiggInfo[k] = num
	}
	return DiggInfo
}

// CommentClear 清空所有文章的评论数
func CommentClear() error {
	err := global.Redis.Del(context.Background(), CommentPrefix).Err()
	return err
}
