package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"time"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 100,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Second)
	defer cancel()
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logrus.Error(err)
		return
	}
}

func main() {
	err := rdb.Set(context.Background(), "name", "xiaoxin", 10*time.Second).Err()
	fmt.Println(err)
	cmd := rdb.Keys(context.Background(), "*")
	keys, err := cmd.Result()
	fmt.Println(keys, err)
}
