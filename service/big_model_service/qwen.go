package big_model_service

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"net/http"
	"strings"
)

type QwenModel struct {
	SessionID uint
}

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

func (qwen QwenModel) Send(content string) (msgChan chan string, err error) {
	set := global.Config.BigModel.Setting
	msgChan = make(chan string, 0)
	baseUrl := "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation"

	req := Request{
		Model: "qwen-turbo",
		Input: Input{},
		Parameters: Parameters{
			IncrementalOutput: true,
		},
	}
	// 查当前这个会话都聊了哪些
	if qwen.SessionID != 0 {

		var sessionModel models.BigModelSessionModel
		err = global.DB.Preload("RoleModel").Take(&sessionModel, qwen.SessionID).Error
		if err != nil {
			return nil, errors.New("不存在的会话")
		}

		req.Input.Messages = append(req.Input.Messages, Message{
			Role:    "system",
			Content: sessionModel.RoleModel.Prompt,
		})

		// 加历史记录
		var chatList []models.BigModelChatModel
		global.DB.Order("created_at asc").Find(&chatList, "session_id = ?", qwen.SessionID)
		for _, model := range chatList {
			req.Input.Messages = append(req.Input.Messages,
				Message{
					Role:    "user",
					Content: model.Content,
				},
				Message{
					Role:    "assistant",
					Content: model.BotContent,
				},
			)
		}
	}

	req.Input.Messages = append(req.Input.Messages, Message{
		Role:    "user",
		Content: content,
	})

	byteData, _ := json.Marshal(req)
	//fmt.Println(string(byteData))
	buf := bytes.NewBuffer(byteData)

	request, err := http.NewRequest("POST", baseUrl, buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("Authorization", "Bearer "+set.ApiKey)
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
