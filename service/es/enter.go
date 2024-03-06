package es

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"strings"
)

type Option struct {
	models.Page
	Fields []string
	Tag    string
}

type SortField struct {
	Field     string
	Ascending bool
}

func (o *Option) GetFrom() int {
	if o.Limit == 0 { // 默认每页10条
		o.Limit = 10
	}
	if o.PageNum == 0 {
		o.PageNum = 1
	}
	return (o.PageNum - 1) * o.Limit
}

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
		return
	}
	err = json.Unmarshal(res.Source, &article)
	if err != nil {
		return
	}
	article.ID = res.Id
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
