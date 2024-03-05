package utils

import (
	"crypto/md5"
	"encoding/hex"
	"gvb_server/global"
	"io"
	"mime/multipart"
	"os"
)

// InList 判断字符串是否在列表中
func InList(key string, list []string) bool {
	for _, v := range list {
		if v == key {
			return true
		}
	}
	return false
}

// MD5 md5加密
func MD5(src []byte) string {
	m := md5.New()
	m.Write(src)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

// GetImageMD5 获取图片的md5值
func GetImageMD5(file *multipart.FileHeader) (byteData []byte, imageHash string, err error) {
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteData, err = io.ReadAll(fileObj)
	if err != nil {
		global.Log.Error(err)
	}
	imageHash = MD5(byteData)
	return
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
