package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"strings"
)

type SearchData struct {
	Body  string `json:"body"`  // 内容
	Slug  string `json:"slug"`  // 跳转地址
	Title string `json:"title"` // 标题（不是文章的大标题，而是一篇文章中的小标题）
}

func main() {
	var data = "## 环境搭建 \n\n 创建一个新的文件夹，然后在文件夹中创建一个新的文件\n\n## 运行程序\n 执行main.go\n```go\n#Linux\nfmt.Print('hello, world')```\n执行结束\n## 拜拜\n jj"
	list := GetSearchIndexDataByContent("/article/dI86s139ER5fd", "测试文章", data) // 这里的title是文章的大标题
	for _, v := range list {
		fmt.Println(v)
	}
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
			Title: v,
			Slug:  id + getSlug(v),
			Body:  bodyList[i]})
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
