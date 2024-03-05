package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
)

type TagsResponse struct {
	Tag           string   `json:"tag"`
	Count         int      `json:"count"`
	ArticleIDList []string `json:"article_id_list"`
}

type TagsType struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count"`
		Articles struct {
			DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int `json:"sum_other_doc_count,omitempty"`
			Buckets                 []struct {
				Key      string `json:"key"`
				DocCount int    `json:"doc_count"`
			} `json:"buckets,omitempty"`
		} `json:"articles"`
	} `json:"buckets"`
}

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接es
	global.ESClient = core.EsConnect()

	// 需要返回的数据形式
	// [{"tag":"go", "article_count": 2, "article_list": [YBJoDo4Beq8OFDNutYS1, YRJpDo4Beq8OFDNulYR4]}]
	agg := elastic.NewTermsAggregation().Field("tags")
	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))
	query := elastic.NewBoolQuery()
	res, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Errorf("查询失败，%s", err.Error())
		return
	}

	var data TagsType
	_ = json.Unmarshal(res.Aggregations["tags"], &data)

	var resList = make([]TagsResponse, 0)
	for _, bucket := range data.Buckets {
		var articleList = make([]string, 0)
		for _, article := range bucket.Articles.Buckets {
			articleList = append(articleList, article.Key)
		}
		resList = append(resList, TagsResponse{
			Tag:           bucket.Key,
			Count:         bucket.DocCount,
			ArticleIDList: articleList,
		})
	}

	fmt.Println(resList)
}
