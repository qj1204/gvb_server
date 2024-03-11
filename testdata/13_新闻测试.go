package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Params struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}

type NewsResponse struct {
	Code int `json:"code"`
	Data []struct {
		Index    string `json:"index"`
		Title    string `json:"title"`
		HotValue string `json:"hotValue"`
		Link     string `json:"link"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func main() {
	params := Params{
		ID:   "mproPpoq6O",
		Size: 100,
	}
	reqParam, _ := json.Marshal(params)
	reqBody := strings.NewReader(string(reqParam))
	httpReq, err := http.NewRequest("POST", "https://api.codelife.cc/api/top/list", reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("signaturekey", "U2FsdGVkX19f0CjpVsmX3u9uWB3jPoruKzCPueFzHjU=")
	httpReq.Header.Add("version", "1.3.52")

	client := http.Client{
		Timeout: 2 * time.Second,
	}
	// DO: HTTP请求
	httpRes, err := client.Do(httpReq)
	if err != nil {
		fmt.Println(err)
		return
	}
	byteData, err := io.ReadAll(httpRes.Body)
	var response NewsResponse
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response.Code != 200 {
		fmt.Println(response.Msg)
		return
	}
	fmt.Println(response.Data)
}
