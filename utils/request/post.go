package request

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Post(url string, params any, header map[string]any, timeout time.Duration) (resp *http.Response, err error) {
	reqParam, _ := json.Marshal(params)
	reqBody := strings.NewReader(string(reqParam))
	httpReq, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return
	}
	httpReq.Header.Add("Content-Type", "application/json")
	for k, v := range header {
		switch val := v.(type) {
		case string:
			httpReq.Header.Add(k, val)
		case int:
			httpReq.Header.Add(k, strconv.Itoa(val))
		}
	}

	client := http.Client{
		Timeout: timeout,
	}
	httpRes, err := client.Do(httpReq)
	return httpRes, err
}
