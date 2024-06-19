package main

import (
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"strings"
)

func main() {
	// github.com/russross/blackfriday
	unsafe := blackfriday.MarkdownCommon([]byte("### 你好\n ```python\nprint('你好')\n```\n - 123 \n \n<script>alert(123)</script>\n\n ![图片](http://xxx.com)"))
	fmt.Println(string(unsafe))
	// github.com/PuerkitoBio/goquery
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	//fmt.Println(doc.Text())
	nodes := doc.Find("h1").Nodes
	fmt.Println(nodes)
	fmt.Println(doc.Text())

	converter := md.NewConverter("", true, nil)
	html, _ := doc.Html()
	markdown, err := converter.ConvertString(html)
	fmt.Println(markdown, err)

}
