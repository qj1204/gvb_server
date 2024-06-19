package cron_service

import (
	"github.com/robfig/cron/v3"
	"time"
)

func CronInit() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	Cron := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))
	Cron.AddFunc("0 0 0 * * *", SyncArticleData) // 每日0点同步文章数据
	Cron.AddFunc("0 0 0 * * *", SyncCommentData) // 每日0点同步评论数据
	Cron.Start()
}
