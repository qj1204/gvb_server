package es

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis"
	"strings"
)

func CommonList(option Option) (list []models.ArticleModel, count int, err error) {
	boolSearch := elastic.NewBoolQuery()
	if option.Key != "" {
		boolSearch.Must(elastic.NewMultiMatchQuery(option.Key, option.Fields...))
	}
	if option.Tag != "" {
		boolSearch.Must(elastic.NewMultiMatchQuery(option.Tag, "tags"))
	}

	sortField := SortField{ // 默认按照创建时间倒序
		Field:     "created_at",
		Ascending: false,
	}
	if option.Sort != "" {
		_list := strings.Split(option.Sort, ",")
		if len(_list) == 2 && (_list[1] == "asc" || _list[1] == "desc") {
			sortField.Field = _list[0]
			if _list[1] == "asc" {
				sortField.Ascending = true
			}
		}
	}

	res, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Highlight(elastic.NewHighlight().Field("title")). // 高亮显示title字段
		From(option.GetFrom()).
		Sort(sortField.Field, sortField.Ascending).
		Size(option.Limit).
		Do(context.Background())
	if err != nil {
		return
	}
	count = int(res.Hits.TotalHits.Value) // 搜索到的结果总条数
	articleList := make([]models.ArticleModel, 0)

	// 从redis中获取点赞数、浏览量
	diggInfo := redis.GetDiggInfo()
	lookInfo := redis.GetLookInfo()

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
		if title, ok := hit.Highlight["title"]; ok { // 如果title字段有高亮显示
			article.Title = title[0]
		}
		article.ID = hit.Id
		// 从redis中获取点赞数、浏览量
		article.DiggCount = article.DiggCount + diggInfo[article.ID]
		article.LookCount = article.LookCount + lookInfo[article.ID]
		if diggInfo[article.ID] == 0 && lookInfo[article.ID] == 0 {
			global.Log.Infof("%s 点赞数和浏览量无变化", article.Title)
			continue
		}

		// 更新文章的点赞数、浏览量
		_, err = global.ESClient.Update().
			Index(models.ArticleModel{}.Index()).
			Id(article.ID).
			Doc(map[string]int{
				"digg_count": article.DiggCount,
				"look_count": article.LookCount,
			}).
			Do(context.Background())
		if err != nil {
			global.Log.Errorf("更新失败，%s", err.Error())
			continue
		}
		global.Log.Infof("%s点赞数同步成功，点赞数%d，浏览量%d", article.Title, article.DiggCount, article.LookCount)
		articleList = append(articleList, article)
	}
	redis.DiggClear()
	redis.LookClear()
	return articleList, count, nil
}

func CommonDetail(id string) (article models.ArticleModel, err error) {
	res, err := global.ESClient.Get().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		return
	}
	err = json.Unmarshal(res.Source, &article)
	if err != nil {
		return
	}
	article.ID = res.Id
	article.DiggCount = article.DiggCount + redis.GetDigg(id)
	article.LookCount = article.LookCount + redis.GetLook(id)
	return
}

func CommonDetailByKeyword(key string) (article models.ArticleModel, err error) {
	res, err := global.ESClient.Search().
		Index(models.ArticleModel{}.Index()).
		Query(elastic.NewTermQuery("keyword", key)).
		Do(context.Background())
	if err != nil {
		return
	}
	if res.Hits.TotalHits.Value == 0 {
		return article, errors.New("文章不存在")
	}
	hit := res.Hits.Hits[0]
	err = json.Unmarshal(hit.Source, &article)
	if err != nil {
		return
	}
	article.ID = hit.Id
	return
}

func ArticleUpdate(id string, maps map[string]any) error {
	_, err := global.ESClient.
		Update().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Doc(maps).
		Do(context.Background())
	return err
}