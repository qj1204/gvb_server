package redis

import (
	"context"
	"gvb_server/global"
	"strconv"
)

type Count struct {
	Index string // 索引
}

// Set 设置某个索引的数值，如果不存在则创建，存在则+1
func (this Count) Set(id string) error {
	num, _ := global.Redis.HGet(context.Background(), this.Index, id).Int()
	num++
	err := global.Redis.HSet(context.Background(), this.Index, id, num).Err()
	return err
}

// SetCount 在原有的基础上增加数值
func (this Count) SetCount(id string, num int) error {
	oldNum, _ := global.Redis.HGet(context.Background(), this.Index, id).Int()
	err := global.Redis.HSet(context.Background(), this.Index, id, oldNum+num).Err()
	return err
}

// Get 获取某个索引的数据
func (this Count) Get(id string) int {
	num, _ := global.Redis.HGet(context.Background(), this.Index, id).Int()
	return num
}

// GetInfo 获取所有索引的数据
func (this Count) GetInfo() map[string]int {
	var DiggInfo = make(map[string]int)
	maps := global.Redis.HGetAll(context.Background(), this.Index).Val()
	for k, v := range maps {
		num, _ := strconv.Atoi(v)
		DiggInfo[k] = num
	}
	return DiggInfo
}

// Delete 删除某个索引的某一条数据
func (this Count) Delete(id string) error {
	err := global.Redis.HDel(context.Background(), this.Index, id).Err()
	return err
}

// Clear 清空所有索引的数据
func (this Count) Clear() error {
	err := global.Redis.Del(context.Background(), this.Index).Err()
	return err
}
