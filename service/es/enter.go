package es

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
)

func CommonList(key string, page int, limit int) (list []models.ArticleModel, count int, err error) {
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
		boolSearch.Must(elastic.NewMatchQuery("title", key))
	}
	if limit == 0 { // 默认每页10条
		limit = 10
	}
	if page == 0 {
		from = 1
	}
	res, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())
	if err != nil {
		global.Log.Errorf("查询失败，%s", err.Error())
		return
	}
	count = int(res.Hits.TotalHits.Value) // 搜索到的结果总条数
	articleList := make([]models.ArticleModel, 0)
	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			global.Log.Error(err.Error())
			continue
		}
		err = json.Unmarshal(data, &article)
		if err != nil {
			global.Log.Error(err.Error())
			continue
		}
		article.ID = hit.Id
		articleList = append(articleList, article)
	}
	return articleList, count, nil
}

func CommonDetail(id string) (article models.ArticleModel, err error) {
	res, err := global.ESClient.Get().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		global.Log.Errorf("查询失败，%s", err.Error())
		return
	}
	err = json.Unmarshal(res.Source, &article)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}
	article.ID = res.Id
	return
}

func CommonDetailByKeyword(key string) (article models.ArticleModel, err error) {
	res, err := global.ESClient.Search().
		Index(models.ArticleModel{}.Index()).
		Query(elastic.NewTermQuery("keyword", key)).
		Size(1).
		Do(context.Background())
	if err != nil {
		global.Log.Errorf("查询失败，%s", err.Error())
		return
	}
	if res.Hits.TotalHits.Value == 0 {
		return article, errors.New("文章不存在")
	}
	hit := res.Hits.Hits[0]
	err = json.Unmarshal(hit.Source, &article)
	if err != nil {
		global.Log.Error(err)
		return
	}
	article.ID = hit.Id
	return
}
