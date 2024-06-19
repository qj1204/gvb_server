package utils

import (
	"crypto/md5"
	"encoding/hex"
	"gvb_server/global"
	"io"
	"mime/multipart"
)

// Md5 md5
func Md5(src []byte) string {
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
	imageHash = Md5(byteData)
	return
}
