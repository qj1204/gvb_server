package article

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"time"
)

type CalendarResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type BucketsType struct {
	Buckets []struct {
		KeyAsString string `json:"key_as_string"`
		Key         int64  `json:"key"`
		DocCount    int    `json:"doc_count"`
	} `json:"buckets"`
}

var DateCount = make(map[string]int)

func (this *ArticleApi) ArticleCalendarView(c *gin.Context) {
	// 时间聚合
	agg := elastic.NewDateHistogramAggregation().Field("created_at").CalendarInterval("day")

	// 时间段搜索，搜索最近一年的数据
	now := time.Now()
	aYearAgo := now.AddDate(-1, 0, 0)
	format := "2006-01-02 15:04:05"
	query := elastic.NewRangeQuery("created_at").
		Gte(aYearAgo.Format(format)).
		Lte(now.Format(format))

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("article_calendar", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("查询失败", c)
		return
	}

	var data BucketsType
	_ = json.Unmarshal(result.Aggregations["article_calendar"], &data)
	// {[{2024-03-04 00:00:00 1709510400000 1} {2024-03-05 00:00:00 1709596800000 1}]}

	var resList = make([]CalendarResponse, 0)
	for _, bucket := range data.Buckets {
		time2, _ := time.Parse(format, bucket.KeyAsString)
		DateCount[time2.Format("2006-01-02")] = bucket.DocCount
	}
	days := int(now.Sub(aYearAgo).Hours() / 24)
	for i := 0; i <= days; i++ {
		day := aYearAgo.AddDate(0, 0, i).Format("2006-01-02")
		resList = append(resList, CalendarResponse{
			Date:  day,
			Count: DateCount[day],
		})
	}
	response.OkWithData(resList, c)
}
