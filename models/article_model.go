package models

import (
	"context"
	"gvb_server/global"
	"gvb_server/models/common/ctype"
)

// ArticleModel 文章表
type ArticleModel struct {
	ID        string `json:"id" structs:"id"`                 // es的id
	CreatedAt string `json:"created_at" structs:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at" structs:"updated_at"` // 更新时间

	Title    string `json:"title" structs:"title"`                // 文章标题
	Keyword  string `json:"keyword,omit(list)" structs:"keyword"` // 文章关键字（在list场景中，过滤掉keyword字段）
	Abstract string `json:"abstract" structs:"abstract"`          // 文章简介
	Content  string `json:"content,omit(list)" structs:"content"` // 文章内容（在list场景中，过滤掉content字段）

	LookCount     int `json:"look_count" structs:"look_count"`         // 文章浏览量
	CommentCount  int `json:"comment_count" structs:"comment_count"`   // 文章评论量
	DiggCount     int `json:"digg_count" structs:"digg_count"`         // 文章点赞量
	CollectsCount int `json:"collects_count" structs:"collects_count"` // 文章收藏量

	// --------用户 一对多 文章--------
	UserID       uint   `json:"user_id" structs:"user_id"`               // 用户ID
	UserNickName string `json:"user_nick_name" structs:"user_nick_name"` // 文章作者（不是冗余，用空间换时间，节省查找的时间）
	UserAvatar   string `json:"user_avatar" structs:"user_avatar"`       // 文章作者头像（不是冗余，用空间换时间，节省查找的时间）

	Category string `json:"category" structs:"category"` // 文章分类
	Source   string `json:"source" structs:"source"`     // 文章来源
	Link     string `json:"link" structs:"link"`         // 原文链接

	// --------文章 belongs to 封面--------
	BannerID  uint   `json:"banner_id" structs:"banner_id"`   // 文章封面ID
	BannerUrl string `json:"banner_url" structs:"banner_url"` // 文章封面（不是冗余，用空间换时间，节省查找的时间）

	Tags ctype.Array `json:"tags" structs:"tags"` // 文章标签
}

func (this ArticleModel) Index() string {
	return "article_index"
}

func (this ArticleModel) Mapping() string { // keyword类型不会被分词（用于精确匹配），text类型会被分词
	return `{
	"settings": {
		"index": {
			"max_result_window": "100000"
		}
	},
	"mappings": {
		"properties": {
			"title":{"type": "text"},
			"keyword":{"type": "keyword"},
			"abstract":{"type": "text"},
			"content":{"type": "text"},
			"look_count":{"type": "integer"},
			"comment_count":{"type": "integer"},
			"digg_count":{"type": "integer"},
			"collects_count":{"type": "integer"},
			"user_id":{"type": "integer"},
			"user_nick_name":{"type": "keyword"},
			"user_avatar":{"type": "keyword"},
			"category":{"type": "keyword"},
			"source":{"type": "keyword"},
			"link":{"type":	"keyword"},
			"banner_id":{"type": "integer"},
			"banner_url":{"type": "keyword"},
			"tags":{"type": "keyword"},
			"created_at": {
				"type": "date",
				"null_value": "null",
				"format": "[yyyy-MM-dd HH:mm:ss]"
			},
			"updated_at": {
				"type": "date",
				"null_value": "null",
				"format": "[yyyy-MM-dd HH:mm:ss]"
			}
		}
	}
}`
}

// IndexExists 判断索引是否存在
func (this ArticleModel) IndexExists() bool {
	exists, err := global.ESClient.
		IndexExists(this.Index()).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err.Error())
	}
	return exists
}

// CreateIndex 创建索引
func (this ArticleModel) CreateIndex() error {
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
func (this ArticleModel) RemoveIndex() error {
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

// InsertArticle 添加文章
func (this *ArticleModel) InsertArticle() error {
	indexResponse, err := global.ESClient.
		Index().
		Index(this.Index()).
		BodyJson(this).
		Do(context.Background())
	if err != nil {
		global.Log.Errorf("添加文章失败，%s", err.Error())
		return err
	}
	global.Log.Infof("添加文章成功，%#v", indexResponse)
	this.ID = indexResponse.Id
	return nil
}
