package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"gvb_server/config"
	"gvb_server/global"
	"time"
)

// 获取上传的token
func getToken(q config.QiNiu) string {
	accessKey := q.AccessKey
	secretKey := q.SecretKey
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

// getCfg 获取上传配置
func getCfg(q config.QiNiu) storage.Config {
	cfg := storage.Config{}
	// 空间对应的机房
	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	return cfg
}

// UploadImage 上传图片
func UploadImage(data []byte, imageName string, prefix string) (filePath string, err error) {
	q := global.Config.QiNiu
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("请配置七牛云的AccessKey和SecretKey")
	}
	//if float64(len(data))/float64(1024*1024) > float64(q.Size) {
	//	return "", errors.New("文件大小超过限制")
	//}
	upToken := getToken(q)
	cfg := getCfg(q)
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))

	// 获取当前时间
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf("%s/%s_%s", prefix, now, imageName)

	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", q.CDN, ret.Key), nil
}
