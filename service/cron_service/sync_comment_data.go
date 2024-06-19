package cron_service

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_service"
)

func SyncCommnetData() {
	commentInfo := redis_service.NewCommentDiggCount().GetInfo()
	for k, v := range commentInfo {
		var comment models.CommentModel
		if err := global.DB.Take(&comment, k); err != nil {
			global.Log.Error(err)
			continue
		}
		if err := global.DB.Model(&comment).Update("digg_count", gorm.Expr("digg_count + ?", v)).Error; err != nil {
			global.Log.Error(err)
			continue
		}
		global.Log.Infof("%s 更新成功，新的点赞数为：%d", comment.ID, comment.DiggCount)
	}
	redis_service.NewCommentDiggCount().Clear()
}
