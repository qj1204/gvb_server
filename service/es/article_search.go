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

	// 从redis中获取点赞数、浏览量、评论数
	diggInfo := redis.NewArticleDiggCount().GetInfo()
	lookInfo := redis.NewArticleLookCount().GetInfo()
	commentInfo := redis.NewArticleCommentCount().GetInfo()

	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err.Error())
			continue
		}
		if title, ok := hit.Highlight["title"]; ok { // 如果title字段有高亮显示
			article.Title = title[0]
		}
		article.ID = hit.Id
		// 从redis中获取点赞数、浏览量、评论数
		article.DiggCount = article.DiggCount + diggInfo[article.ID]
		article.LookCount = article.LookCount + lookInfo[article.ID]
		article.CommentCount = article.CommentCount + commentInfo[article.ID]

		articleList = append(articleList, article)
	}
	return articleList, count, nil
}

func CommonDetail(id string) (article models.ArticleModel, err error) {
	result, err := SearchArticleByID(id)
	if err != nil {
		return
	}
	err = json.Unmarshal(result.Source, &article)
	if err != nil {
		return
	}
	article.ID = result.Id
	article.DiggCount = article.DiggCount + redis.NewArticleDiggCount().Get(id)
	article.LookCount = article.LookCount + redis.NewArticleLookCount().Get(id)
	article.CommentCount = article.CommentCount + redis.NewArticleCommentCount().Get(id)
	return
}

func CommonDetailByKeyword(key string) (article models.ArticleModel, err error) {
	result, err := SearchArticleByTitle(key)
	if err != nil {
		return
	}
	if result.Hits.TotalHits.Value == 0 {
		return article, errors.New("文章不存在")
	}
	hit := result.Hits.Hits[0]
	err = json.Unmarshal(hit.Source, &article)
	if err != nil {
		return
	}
	article.ID = hit.Id
	return
}

func SearchArticleByID(id string) (*elastic.GetResult, error) {
	result, err := global.ESClient.Get().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Do(context.Background())
	return result, err
}

func SearchArticleByTitle(title string) (*elastic.SearchResult, error) {
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewTermQuery("keyword", title)).
		Do(context.Background())
	return result, err
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

func CommonIDListByTag(tag string) (list []string, err error) {
	sortField := SortField{ // 默认按照创建时间倒序
		Field:     "created_at",
		Ascending: false,
	}

	res, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMultiMatchQuery(tag, "tags")).
		Sort(sortField.Field, sortField.Ascending).
		Size(1000).
		Do(context.Background())
	if err != nil {
		return
	}
	IDList := make([]string, 0)

	for _, hit := range res.Hits.Hits {
		IDList = append(IDList, hit.Id)
	}
	return IDList, nil
}
