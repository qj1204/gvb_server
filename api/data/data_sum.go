package data

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
)

type DataSumResponse struct {
	UserCount       int64 `json:"user_count"`
	ArticleCount    int64 `json:"article_count"`
	MessageCount    int64 `json:"message_count"`
	ChatGroupCount  int64 `json:"chat_group_count"`
	TodayLoginCount int64 `json:"today_login_count"`
	TodaySignCount  int64 `json:"today_sign_count"`
}

func (this *DataApi) DataSumView(c *gin.Context) {
	var userCount, articleCount, messageCount, chatGroupCount int64
	var todayLoginCount, todaySignCount int64

	global.DB.Model(models.UserModel{}).Count(&userCount)
	global.DB.Model(models.MessageModel{}).Count(&messageCount)
	global.DB.Model(models.ChatModel{IsGroup: true}).Count(&chatGroupCount)
	res, _ := global.ESClient.Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Do(context.Background())
	articleCount = res.TotalHits()

	global.DB.Model(models.LoginDataModel{}).Where("to_days(created_at)=to_days(now())").Count(&todayLoginCount)
	global.DB.Model(models.UserModel{}).Where("to_days(created_at)=to_days(now())").Count(&todaySignCount)

	response.OkWithData(DataSumResponse{
		UserCount:       userCount,
		ArticleCount:    articleCount,
		MessageCount:    messageCount,
		ChatGroupCount:  chatGroupCount,
		TodayLoginCount: todayLoginCount,
		TodaySignCount:  todaySignCount,
	}, c)
}
