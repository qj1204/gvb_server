package article

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
)

type TagsResponse struct {
	Tag           string   `json:"tag"`
	Count         int      `json:"count"`
	ArticleIDList []string `json:"article_id_list"`
	CreatedAt     string   `json:"created_at"`
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

// ArticleTagListView 标签文章列表
// @Tags 文章管理
// @Summary 标签文章列表
// @Description 标签文章列表
// @Param data query models.Page   false  "表示多个参数"
// @Router /api/articles/tags [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[TagsResponse]}
func (ArticleApi) ArticleTagListView(c *gin.Context) {
	// 需要返回的数据形式
	// [{"tag":"go", "article_count": 2, "article_list": [YBJoDo4Beq8OFDNutYS1, YRJpDo4Beq8OFDNulYR4]}]

	var cr models.Page
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	// 查询标签总数
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Aggregation("tags", elastic.NewCardinalityAggregation().Field("tags")).
		Do(context.Background())
	cTag, _ := result.Aggregations.Cardinality("tags")
	count := int64(*cTag.Value)

	// 标签聚合分页
	if cr.Limit == 0 {
		cr.Limit = 10
	}
	offset := (cr.PageNum - 1) * cr.Limit
	offset = max(offset, 0)
	agg := elastic.NewTermsAggregation().Field("tags")
	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))
	agg.SubAggregation("page", elastic.NewBucketSortAggregation().From(offset).Size(cr.Limit))

	result, err = global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewBoolQuery()).
		Aggregation("tags", agg).
		Do(context.Background())
	if err != nil {
		global.Log.Errorf("查询失败，%s", err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}

	var data TagsType
	_ = json.Unmarshal(result.Aggregations["tags"], &data)

	var tagTitleList = make([]string, 0)

	// 组装返回数据
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
		tagTitleList = append(tagTitleList, bucket.Key)
	}

	// 在mysql中查询标签的创建时间，并添加到返回数据中
	var tagList []models.TagModel
	global.DB.Find(&tagList, "title in ?", tagTitleList)
	var tagMap = make(map[string]string)
	for _, tag := range tagList {
		tagMap[tag.Title] = tag.CreatedAt.Format("2006-01-02 15:04:05")
	}
	// 不能用for range，因为for range是值拷贝，修改值不会影响原值
	for i := 0; i < len(resList); i++ {
		resList[i].CreatedAt = tagMap[resList[i].Tag]
	}
	response.OkWithList(resList, count, c)
}
