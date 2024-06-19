package big_model_service

import (
	"errors"
	"gvb_server/global"
)

type BigModelInterface interface {
	Send(content string) (msgChan chan string, err error)
}

func Send(sessionID uint, content string) (msgChan chan string, err error) {
	var ser BigModelInterface
	switch global.Config.BigModel.Setting.Name {
	case "qwen":
		ser = QwenModel{SessionID: sessionID}
	case "wenxin":
	case "xinghuo":
	case "tiangong":
	case "ChatGPT":
	default:
		return nil, errors.New("不支持的大模型")
	}
	return ser.Send(content)
}
