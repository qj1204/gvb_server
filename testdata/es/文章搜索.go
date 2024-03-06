package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.ESClient = core.EsConnect()
	res, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMultiMatchQuery("java", "title", "abstract", "content")).
		Highlight(elastic.NewHighlight().Field("title")). // 高亮显示title字段
		Size(10).
		Do(context.Background())
	if err != nil {
		global.Log.Errorf("查询失败，%s", err.Error())
		return
	}
	for _, hit := range res.Hits.Hits {
		fmt.Println(string(hit.Source))
		fmt.Println(hit.Highlight)
	}
}
