package redis_service

import (
	"context"
	"gvb_server/global"
	"strconv"
)

type CountDB struct {
	Index string // 索引
}

// Set 设置某个索引的数值，如果不存在则创建，存在则+1
func (c CountDB) Set(id string) error {
	num, _ := global.Redis.HGet(context.Background(), c.Index, id).Int()
	num++
	err := global.Redis.HSet(context.Background(), c.Index, id, num).Err()
	return err
}

// SetCount 在原有的基础上增加数值
func (c CountDB) SetCount(id string, num int) error {
	oldNum, _ := global.Redis.HGet(context.Background(), c.Index, id).Int()
	newNum := oldNum + num
	err := global.Redis.HSet(context.Background(), c.Index, id, newNum).Err()
	return err
}

// Get 获取某个索引的数据
func (c CountDB) Get(id string) int {
	num, _ := global.Redis.HGet(context.Background(), c.Index, id).Int()
	return num
}

// GetInfo 获取所有索引的数据
func (c CountDB) GetInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := global.Redis.HGetAll(context.Background(), c.Index).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo
}

// Delete 删除某个索引的某一条数据
func (c CountDB) Delete(id string) error {
	err := global.Redis.HDel(context.Background(), c.Index, id).Err()
	return err
}

// Clear 清空所有索引的数据
func (c CountDB) Clear() {
	global.Redis.Del(context.Background(), c.Index)
}
