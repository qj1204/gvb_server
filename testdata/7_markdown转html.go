package main

import (
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"strings"
)

func main() {
	// markdown转html
	//unsafe := blackfriday.MarkdownCommon([]byte("## 你好\n ```go\nfmt.Println('hello')\n```\n - 123 \n <script>alert</script>\n\n ![图片](http://xx.com)"))
	unsafe := blackfriday.MarkdownCommon([]byte("## 二级标题\n 哈哈哈哈哈 我是go语言进阶\n<script>alert(123)</script>"))
	fmt.Println(string(unsafe))
	// html获取文本内容，xss过滤
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	//fmt.Println("原文本：")
	//fmt.Println(doc.Text())
	doc.Find("script").Remove()
	//fmt.Println("去掉script标签的文本：")
	//fmt.Println(doc.Text())

	//html转markdown
	converter := md.NewConverter("", true, nil)
	html, _ := doc.Html()
	markdown, _ := converter.ConvertString(html)
	fmt.Println("html转markdown：")
	fmt.Println(markdown)
}
