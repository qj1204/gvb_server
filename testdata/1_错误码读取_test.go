package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gvb_server/models/common/response"
	"os"
	"testing"
)

// 如果错误码是以json格式存储在文件中，可以使用以下方法读取

const FILE = "models/common/response/err_code.json"

type ErrorMap map[response.ErrCode]string

func TestErrCode(t *testing.T) {
	byteData, err := os.ReadFile(FILE)
	if err != nil {
		logrus.Error("读取错误码文件失败: ", err)
		return
	}
	var errorMap ErrorMap
	err = json.Unmarshal(byteData, &errorMap)
	if err != nil {
		logrus.Error("解析错误码文件失败: ", err)
		return
	}
	fmt.Println(errorMap[response.SettingError])
	fmt.Println(errorMap[response.ArgumentError])
}
