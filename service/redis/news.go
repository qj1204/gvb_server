package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"time"
)

const NewsIndex = "news"

// SetNews 设置某个索引的数值
func SetNews(key string, newsData []models.NewsData) error {
	byteData, _ := json.Marshal(newsData)
	err := global.Redis.Set(context.Background(), fmt.Sprintf("%s_%s", NewsIndex, key), byteData, 10*time.Minute).Err()
	return err
}

// GetNews  获取某个索引的数据
func GetNews(key string) (newsData []models.NewsData, err error) {
	res := global.Redis.Get(context.Background(), fmt.Sprintf("%s_%s", NewsIndex, key)).Val()
	err = json.Unmarshal([]byte(res), &newsData)
	return
}
