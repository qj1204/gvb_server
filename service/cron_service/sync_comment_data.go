package cron_service

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_service"
)

// SyncCommentData 同步评论数据到数据库
func SyncCommentData() {
	commentDiggInfo := redis_service.NewCommentDiggCount().GetInfo()
	for key, count := range commentDiggInfo {
		var comment models.CommentModel
		err := global.DB.Take(&comment, key).Error
		if err != nil {
			global.Log.Error(err)
			continue
		}
		err = global.DB.Model(&comment).Update("digg_count", gorm.Expr("digg_count + ?", count)).Error
		if err != nil {
			global.Log.Error(err)
			continue
		}
		global.Log.Infof("%s 更新成功 新的点赞数为：%d", comment.Content, comment.DiggCount)
	}
	redis_service.NewCommentDiggCount().Clear()
}
