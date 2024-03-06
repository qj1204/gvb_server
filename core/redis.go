package core

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gvb_server/global"
	"time"
)

func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)
}

func ConnectRedisDB(db int) *redis.Client {
	redisConf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       db,
		PoolSize: redisConf.PoolSize,
	})
	_, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		global.Log.Errorf("redis连接失败 %s", redisConf.Addr())
		return nil
	}
	return rdb
}
