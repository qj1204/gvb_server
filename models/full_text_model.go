package models

import (
	"context"
	"gvb_server/global"
)

type FullTextModel struct {
	ID        string `json:"id" structs:"id"`                 // es的id
	Title     string `json:"title" structs:"title"`           // 文章标题
	ArticleID string `json:"article_id" structs:"article_id"` // 关联的文章id
	Body      string `json:"body" structs:"body"`             // 文章内容
	Slug      string `json:"slug" structs:"slug"`             // 跳转地址
}

func (this FullTextModel) Index() string {
	return "full_text_index"
}

func (this FullTextModel) Mapping() string { // keyword类型不会被分词（用于精确匹配），text类型会被分词
	return `{
	"settings": {
		"index": {
			"max_result_window": "100000"
		}
	},
	"mappings": {
		"properties": {
			"title":{"type": "text"},
			"article_id":{"type": "keyword"},
			"body":{"type": "text"},
			"slug":{"type": "keyword"}
		}
	}
}`
}

// IndexExists 判断索引是否存在
func (this FullTextModel) IndexExists() bool {
	exists, err := global.ESClient.
		IndexExists(this.Index()).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err.Error())
	}
	return exists
}

// CreateIndex 创建索引
func (this FullTextModel) CreateIndex() error {
	if this.IndexExists() {
		// 索引已经存在，删除索引
		this.RemoveIndex()
	}
	// 没有索引，创建索引
	createIndex, err := global.ESClient.
		CreateIndex(this.Index()).
		BodyString(this.Mapping()).
		Do(context.Background())
	if err != nil {
		global.Log.Errorf("创建索引失败, %s", err.Error())
		return err
	}
	if !createIndex.Acknowledged {
		global.Log.Errorf("创建索引失败, %s", err.Error())
		return err
	}
	global.Log.Infof("%s 创建索引成功", this.Index())
	return nil
}

// RemoveIndex 删除索引
func (this FullTextModel) RemoveIndex() error {
	global.Log.Info("索引存在，删除索引")
	indexDelete, err := global.ESClient.
		DeleteIndex(this.Index()).
		Do(context.Background())
	if err != nil {
		global.Log.Errorf("删除索引失败, %s", err.Error())
		return err
	}
	if !indexDelete.Acknowledged {
		global.Log.Errorf("删除索引失败, %s", err.Error())
		return err
	}
	global.Log.Info("删除索引成功")
	return nil
}
