package es_service

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic/v7"
	"github.com/russross/blackfriday"
	"gvb_server/global"
	"gvb_server/models"
	"strings"
)

type SearchData struct {
	Body      string `json:"body"`       // 内容
	Slug      string `json:"slug"`       // 跳转地址
	Title     string `json:"title"`      // 标题（不是文章的大标题，而是一篇文章中的小标题）
	ArticleID string `json:"article_id"` // 关联的文章id
}

func GetSearchIndexDataByContent(id string, title string, content string) (SearchDataList []SearchData) {
	dataList := strings.Split(content, "\n")

	var headList, bodyList []string
	var body string
	var isCode = false
	headList = append(headList, getHeader(title))
	for _, v := range dataList {
		if strings.HasPrefix(v, "```") {
			isCode = !isCode
		}
		if strings.HasPrefix(v, "#") && !isCode {
			headList = append(headList, getHeader(v))
			bodyList = append(bodyList, getBody(body))
			body = ""
			continue
		}
		body += v
		if strings.HasSuffix(v, "```") {
			isCode = !isCode
		}
	}
	bodyList = append(bodyList, getBody(body))

	for i, v := range headList {
		SearchDataList = append(SearchDataList, SearchData{
			Title:     v,
			Slug:      id + getSlug(v),
			Body:      bodyList[i],
			ArticleID: id,
		})
	}
	return
}

func getHeader(head string) string {
	head = strings.ReplaceAll(head, "#", "")
	head = strings.ReplaceAll(head, " ", "")
	return head
}

func getBody(body string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(body))
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	return doc.Text()
}

func getSlug(slug string) string {
	return "#" + slug
}

// AsyncArticleByFullText 同步文章数据到全文搜索索引
func AsyncArticleByFullText(id, title, content string) {
	indexList := GetSearchIndexDataByContent(id, title, content)

	// es批量添加数据
	bulk := global.ESClient.Bulk()
	for _, indexData := range indexList {
		req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(indexData)
		bulk.Add(req)
	}
	result, err := bulk.Do(context.Background())
	if err != nil {
		global.Log.Error(err.Error())

		return
	}
	global.Log.Infof(fmt.Sprintf("%s添加成功，共%d条", title, len(result.Succeeded())))
}

// DeleteFullTextByArticleID 根据文章id删除全文搜索
func DeleteFullTextByArticleID(id string) {
	boolSearch := elastic.NewTermQuery("article_id", id)
	res, _ := global.ESClient.
		DeleteByQuery().
		Index(models.FullTextModel{}.Index()).
		Query(boolSearch).
		Do(context.Background())
	global.Log.Infof("删除文章id为%s的全文搜索数据，共%d条", id, res.Deleted)
}
