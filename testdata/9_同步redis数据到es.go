package main

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.Redis = core.ConnectRedis()
	global.ESClient = core.EsConnect()

	res, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(1000).
		Do(context.Background())
	if err != nil {
		global.Log.Errorf("查询失败，%s", err.Error())
		return
	}

	diggInfo := redis.NewArticleDiggCount().GetInfo()

	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel
		_ = json.Unmarshal(hit.Source, &article)
		newDigg := article.DiggCount + diggInfo[hit.Id]
		if article.DiggCount == newDigg {
			global.Log.Infof("%s 点赞数无变化", article.Title)
			continue
		}
		_, err := global.ESClient.Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"digg_count": newDigg,
			}).
			Do(context.Background())
		if err != nil {
			global.Log.Errorf("更新点赞数失败，%s", err.Error())
			continue
		}
		global.Log.Infof("%s 点赞数同步成功，点赞数%d", article.Title, newDigg)
	}

	global.Log.Info("更新索引成功")
	redis.NewArticleDiggCount().Clear()
}
