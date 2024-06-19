package utils

import (
	"gvb_server/global"
	"os"
)

// InList 判断key是否存在与列表中
func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

// Mkdir 判断路径是否存在，不存在就创建
func Mkdir(path string) {
	_, err := os.ReadDir(path)
	if err != nil {
		// 不存在就创建
		err = os.MkdirAll(path, os.ModePerm) // 递归创建
		if err != nil {
			global.Log.Error(err)
		}
	}
}
