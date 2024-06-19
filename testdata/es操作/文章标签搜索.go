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
			SumOtherDocCount        int `json:"sum_other_doc_count"`
			Buckets                 []struct {
				Key      string `json:"key"`
				DocCount int    `json:"doc_count"`
			} `json:"buckets"`
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

	/*
		[{"tag": "python", "article_count": 2, "article_list": []}]
	*/

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Aggregation("tags", elastic.NewValueCountAggregation().Field("tags")).
		Size(0).
		Do(context.Background())
	cTag, _ := result.Aggregations.Cardinality("tags")
	count := int64(*cTag.Value)
	fmt.Println(count)

	//agg.SubAggregation("article_id", elastic.NewTermsAggregation().Field("_id"))
	//agg.SubAggregation("article_key", elastic.NewTermsAggregation().Field("keyword"))
	//agg.SubAggregation("page", elastic.NewBucketSortAggregation())

	result, err = global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Aggregation("tags", elastic.NewValueCountAggregation().Field("tags")).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		return
	}
	var tagType TagsType
	var tagList = make([]TagsResponse, 0)
	fmt.Println(string(result.Aggregations["tags"]))
	x, _ := result.Aggregations.Cardinality("tags")
	fmt.Println(*x.Value)
	_ = json.Unmarshal(result.Aggregations["tags"], &tagType)
	for _, bucket := range tagType.Buckets {

		var articleList []string
		for _, s := range bucket.Articles.Buckets {
			articleList = append(articleList, s.Key)
		}

		tagList = append(tagList, TagsResponse{
			Tag:           bucket.Key,
			Count:         bucket.DocCount,
			ArticleIDList: articleList,
		})
	}

}
