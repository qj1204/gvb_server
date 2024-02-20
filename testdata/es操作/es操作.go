package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/core"
)

var client *elastic.Client

func EsConnect() *elastic.Client {
	var err error
	sniffOpt := elastic.SetSniff(false)
	host := "http://127.0.0.1:9200"
	c, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""))
	if err != nil {
		logrus.Fatalf("es 连接失败 %s", err.Error())
	}
	return c
}

func init() {
	core.InitConf()
	core.InitLogger()
	client = EsConnect()
}

type DemoModel struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

func (this DemoModel) Index() string {
	return "demo_index"
}

// Create 创建
func Create(data *DemoModel) (err error) {
	indexResponse, err := client.Index().Index(data.Index()).BodyJson(data).Do(context.Background())
	if err != nil {
		logrus.Errorf("添加索引失败，%s", err.Error())
		return err
	}
	logrus.Infof("添加索引成功，%#v", indexResponse)
	data.ID = indexResponse.Id
	return nil
}

// FindList 查询索引
func FindList(key string, page int, limit int) (demoList []DemoModel, count int) {
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
		boolSearch.Must(elastic.NewMatchQuery("title", key))
	}
	if limit == 0 { // 默认每页10条
		limit = 10
	}
	if from == 0 { // 默认第一页
		from = 1
	}
	res, err := client.Search(DemoModel{}.Index()).Query(boolSearch).From((from - 1) * limit).Size(limit).
		Do(context.Background())
	if err != nil {
		logrus.Errorf("查询失败, %s", err.Error())
		return
	}
	count = int(res.Hits.TotalHits.Value) // 搜索到的结果总条数

	for _, hit := range res.Hits.Hits {
		var demo DemoModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		err = json.Unmarshal(data, &demo)
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		demo.ID = hit.Id
		demoList = append(demoList, demo)
	}
	return demoList, count
}

// FindSouceList 查询索引
func FindSouceList(key string, page int, limit int) {
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
		boolSearch.Must(elastic.NewMatchQuery("title", key))
	}
	if limit == 0 { // 默认每页10条
		limit = 10
	}
	if from == 0 { // 默认第一页
		from = 1
	}
	res, err := client.
		Search(DemoModel{}.Index()).
		Query(boolSearch).
		Source(`{"_source": ["title"]}`). // 只返回_source中的某个字段
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())
	if err != nil {
		logrus.Errorf("查询失败, %s", err.Error())
		return
	}
	count := int(res.Hits.TotalHits.Value) // 搜索到的结果总条数
	var demoList []DemoModel
	for _, hit := range res.Hits.Hits {
		var demo DemoModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		err = json.Unmarshal(data, &demo)
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		demo.ID = hit.Id
		demoList = append(demoList, demo)
	}
	fmt.Println(demoList, count)
}

// Update 更新索引
func Update(id string, data *DemoModel) error {
	_, err := client.
		Update().
		Index(data.Index()).
		Id(id).
		Doc(map[string]string{
			"title": data.Title,
		}).
		Do(context.Background())
	if err != nil {
		logrus.Errorf("更新索引失败，%s", err.Error())
		return err
	}
	logrus.Info("更新索引成功")
	return nil
}

// Remove 批量删除索引
func Remove(idList []string) (count int, err error) {
	bulkService := client.Bulk().Index(DemoModel{}.Index()).Refresh("true")
	for _, id := range idList {
		req := elastic.NewBulkDeleteRequest().Index(DemoModel{}.Index()).Id(id)
		bulkService.Add(req)
	}
	res, err := bulkService.Do(context.Background())
	if err != nil {
		logrus.Errorf("删除索引失败，%s", err.Error())
		return 0, err
	}
	logrus.Info("删除索引成功")
	return len(res.Succeeded()), nil
}

func main() {
	//DemoModel{}.CreateIndex()

	//Create(&DemoModel{
	//	Title:     "go基础",
	//	UserID:    1,
	//	CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	//})

	//Create(&DemoModel{
	//	Title:     "go进阶",
	//	UserID:    2,
	//	CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	//})

	//list, count := FindList("go", 1, 10)
	//fmt.Println(list, count)

	// FindSouceList("基础", 1, 10)	// 搜索失效了
	//Update("cT5HxI0BDjsaIRYC8Prk", &DemoModel{Title: "go基础学习"})

	count, err := Remove([]string{"cj5TxI0BDjsaIRYCn_o6"})
	fmt.Println(count, err)
}
