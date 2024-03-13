package cron

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis"
)

func SyncArticleData() {
	// 查询es中的全部数据
	result, err := global.ESClient.Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(1000).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err.Error())
		return
	}
	// 拿到redis中的文章数据
	diggInfo := redis.NewArticleDiggCount().GetInfo()
	lookInfo := redis.NewArticleLookCount().GetInfo()
	commentInfo := redis.NewArticleCommentCount().GetInfo()

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err.Error())
			continue
		}

		if diggInfo[hit.Id] == 0 && lookInfo[hit.Id] == 0 && commentInfo[hit.Id] == 0 {
			global.Log.Infof("%s 无变化", article.Title)
			continue
		}
		article.DiggCount = article.DiggCount + diggInfo[article.ID]
		article.LookCount = article.LookCount + lookInfo[article.ID]
		article.CommentCount = article.CommentCount + commentInfo[article.ID]
		_, err = global.ESClient.Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"digg_count":    article.DiggCount,
				"look_count":    article.LookCount,
				"comment_count": article.CommentCount,
			}).
			Do(context.Background())
		if err != nil {
			global.Log.Errorf("更新失败，%s", err.Error())
			continue
		}
		global.Log.Infof("%s 同步成功，点赞数%d，浏览量%d，评论数%d", article.Title, article.DiggCount, article.LookCount, article.CommentCount)
	}
	// 清空redis中的数据
	redis.NewArticleDiggCount().Clear()
	redis.NewArticleLookCount().Clear()
	redis.NewArticleCommentCount().Clear()
}
