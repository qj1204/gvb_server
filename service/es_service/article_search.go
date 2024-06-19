package es_service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_service"
	"strings"
)

func CommonList(option Option) (list []models.ArticleModel, count int, err error) {
	if option.Query == nil {
		option.Query = elastic.NewBoolQuery()
	}

	if option.Key != "" {
		option.Query.Must(
			elastic.NewMultiMatchQuery(option.Key, option.Fields...),
		)
	}
	if option.Category != "" {
		option.Query.Must(
			elastic.NewMultiMatchQuery(option.Category, "category"),
		)
	}
	if option.Tag != "" {
		option.Query.Must(
			elastic.NewMultiMatchQuery(option.Tag, "tags"),
		)
	}

	sortField := SortField{
		Field:     "created_at",
		Ascending: false, // 从小到大  从大到小
	}
	if option.Sort != "" {
		_list := strings.Split(option.Sort, " ")
		if len(_list) == 2 && (_list[1] == "desc" || _list[1] == "asc") {
			sortField.Field = _list[0]
			if _list[1] == "asc" {
				sortField.Ascending = true
			}
		}
	}

	//data, _ := option.Query.Source()
	//byteData, _ := json.Marshal(data)
	//fmt.Println(string(byteData))

	res, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(option.Query).
		Highlight(elastic.NewHighlight().Field("title")).
		From(option.GetForm()).
		Sort(sortField.Field, sortField.Ascending).
		Size(option.Limit).
		Do(context.Background())
	if err != nil {
		return
	}
	count = int(res.Hits.TotalHits.Value) //搜索到结果总条数
	demoList := make([]models.ArticleModel, 0)

	diggInfo := redis_service.NewArticleDiggCount().GetInfo()
	lookInfo := redis_service.NewArticleLookCount().GetInfo()
	commentInfo := redis_service.NewArticleCommentCount().GetInfo()
	for _, hit := range res.Hits.Hits {
		var model models.ArticleModel
		err = json.Unmarshal(hit.Source, &model)
		if err != nil {
			logrus.Error(err)
			continue
		}
		title, ok := hit.Highlight["title"]
		if ok {
			model.Title = title[0]
		}

		model.ID = hit.Id
		model.DiggCount = model.DiggCount + diggInfo[hit.Id]
		model.LookCount = model.LookCount + lookInfo[hit.Id]
		model.CommentCount = model.CommentCount + commentInfo[hit.Id]

		demoList = append(demoList, model)
	}
	return demoList, count, err
}

func CommonDetail(id string) (model models.ArticleModel, err error) {
	res, err := global.ESClient.Get().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		return
	}
	err = json.Unmarshal(res.Source, &model)
	if err != nil {
		return
	}
	model.ID = res.Id
	model.DiggCount = model.DiggCount + redis_service.NewArticleDiggCount().Get(res.Id)
	model.LookCount = model.LookCount + redis_service.NewArticleLookCount().Get(res.Id)
	model.CommentCount = model.CommentCount + redis_service.NewArticleCommentCount().Get(res.Id)
	return
}

func CommonDetailByKeyword(key string) (model models.ArticleModel, err error) {
	res, err := global.ESClient.Search().
		Index(models.ArticleModel{}.Index()).
		Query(elastic.NewTermQuery("keyword", key)).
		Size(1).
		Do(context.Background())
	if err != nil {
		return
	}
	if res.Hits.TotalHits.Value == 0 {
		return model, errors.New("文章不存在")
	}
	hit := res.Hits.Hits[0]
	err = json.Unmarshal(hit.Source, &model)
	if err != nil {
		return
	}
	model.ID = hit.Id
	return
}

func ArticleUpdate(id string, data map[string]any) error {
	_, err := global.ESClient.
		Update().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Doc(data).Refresh("true").
		Do(context.Background())
	return err
}
