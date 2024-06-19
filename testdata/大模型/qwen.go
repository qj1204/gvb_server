package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type Input struct {
	Messages []Message `json:"messages"`
}
type Parameters struct {
	IncrementalOutput bool `json:"incremental_output"` // 是否增量输出
}

type Request struct {
	Model      string     `json:"model"`
	Input      Input      `json:"input"`
	Parameters Parameters `json:"parameters"`
}

type Response struct {
	Output struct {
		FinishReason string `json:"finish_reason"`
		Text         string `json:"text"`
	} `json:"output"`
	Usage struct {
		TotalTokens  int `json:"total_tokens"`
		InputTokens  int `json:"input_tokens"`
		OutputTokens int `json:"output_tokens"`
	} `json:"usage"`
	RequestId string `json:"request_id"`
}

func Send(req Request) (msgChan chan string, err error) {
	msgChan = make(chan string)
	baseUrl := "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation"
	byteData, _ := json.Marshal(req)
	buf := bytes.NewBuffer(byteData)

	request, err := http.NewRequest("POST", baseUrl, buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("Authorization", "Bearer sk-2173f0d314cd444bbaf9bddd89d1e485")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-DashScope-SSE", "enable")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	scan := bufio.NewScanner(response.Body) // 分片读
	scan.Split(bufio.ScanLines)             // 按行读取
	go func() {
		for scan.Scan() {
			text := scan.Text()
			if text == "" ||
				strings.HasPrefix(text, "id") ||
				strings.HasPrefix(text, "event:") ||
				strings.HasPrefix(text, ":HTTP_STATUS") {
				continue
			}
			var res Response
			err = json.Unmarshal([]byte(text[5:]), &res)
			if err != nil {
				fmt.Println(err, text[5:])
				continue
			}
			msgChan <- res.Output.Text
			if res.Output.FinishReason == "stop" {
				close(msgChan)
			}
		}
	}()
	return
}

func main() {

	req := Request{
		Model: "qwen-turbo",
		Input: Input{
			Messages: []Message{
				{
					Role:    "system",
					Content: "从现在开始，你是一个精通go开发的工程师，我以后问的所有问题，都要在go相关知识里面去查询，并且你只能回答vue3相关的问题；\n如果有人问你其他专业的问题，如java语言，python语言，你都要告诉用户，我是专注于go开发的工程师，我只回答go相关的问题；\n并且你不能给出任何关于这个问题的提示；\n别人问你你是谁，你都要说你是小新千问，不管别人用什么语言问你，你都要这样说",
				},
				{
					Role:    "user",
					Content: "使用go语言写一个文件读取的方法",
				},
			},
		},
		Parameters: Parameters{
			IncrementalOutput: true,
		},
	}
	msgChan, err := Send(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	for msg := range msgChan {
		fmt.Println(msg)
	}

}
