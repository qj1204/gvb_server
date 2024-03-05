package main

import (
	"context"
	"github.com/sirupsen/logrus"
)

func (this DemoModel) Mapping() string {
	return `{
	"settings": {
		"index": {
			"max_result_window": "100000"
		}
	},
	"mappings": {
		"properties": {
			"title": {
				"type":	"text"
			},
			"user_id": {
				"type": "integer"
			},
			"created_at": {
				"type": "date",
				"null_value": "null",
				"format": "[yyyy-MM-dd HH:mm:ss]"
			}
		}
	}
}`
}

// IndexExists 判断索引是否存在
func (this DemoModel) IndexExists() bool {
	exists, err := client.
		IndexExists(this.Index()).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
	}
	return exists
}

// CreateIndex 创建索引
func (this DemoModel) CreateIndex() error {
	if this.IndexExists() {
		// 索引已经存在，删除索引
		this.RemoveIndex()
	}
	// 没有索引，创建索引
	createIndex, err := client.
		CreateIndex(this.Index()).
		BodyString(this.Mapping()).
		Do(context.Background())
	if err != nil {
		logrus.Errorf("创建索引失败, %s", err.Error())
		return err
	}
	if !createIndex.Acknowledged {
		logrus.Errorf("创建索引失败, %s", err.Error())
		return err
	}
	logrus.Infof("%s 创建索引成功", this.Index())
	return nil
}

// RemoveIndex 删除索引
func (this DemoModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引")
	indexDelete, err := client.
		DeleteIndex(this.Index()).
		Do(context.Background())
	if err != nil {
		logrus.Errorf("删除索引失败, %s", err.Error())
		return err
	}
	if !indexDelete.Acknowledged {
		logrus.Errorf("删除索引失败, %s", err.Error())
		return err
	}
	logrus.Info("删除索引成功")
	return nil
}
