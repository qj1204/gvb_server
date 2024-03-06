package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/service/redis"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.Redis = core.ConnectRedis()

	//err := global.Redis.Set(context.Background(), "name", "xiaoxin", 10*time.Second).Err()
	//fmt.Println(err)
	//cmd := global.Redis.Keys(context.Background(), "*")
	//keys, err := cmd.Result()
	//fmt.Println(keys, err)

	redis.Digg("ZhKLDo4Beq8OFDNuzYQB")
	fmt.Println(redis.GetDiggInfo())
	//redis.DiggClear()
}
