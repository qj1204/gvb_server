package redis_service

import (
	"context"
	"encoding/json"
	"fmt"
	"gvb_server/global"
	"time"
)

const newsIndex = "news_index"

type NewData = any

// SetNews 设置某一个数据，重复执行，重复累加
func SetNews(key string, newData []NewData) error {
	byteData, _ := json.Marshal(newData)
	err := global.Redis.Set(context.Background(), fmt.Sprintf("%s_%s", newsIndex, key), byteData, 1*time.Hour).Err()
	return err
}

func GetNews(key string) (newData []NewData, err error) {
	res := global.Redis.Get(context.Background(), fmt.Sprintf("%s_%s", newsIndex, key)).Val()
	err = json.Unmarshal([]byte(res), &newData)
	return
}
