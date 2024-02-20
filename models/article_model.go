package models

import (
	"context"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models/common/ctype"
)

// ArticleModel 文章表
type ArticleModel struct {
	ID        uint   `json:"id"`         // es的id
	CreatedAt string `json:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at"` // 更新时间

	Title    string `json:"title"`    // 文章标题
	Abstract string `json:"abstract"` // 文章简介
	Content  string `json:"content"`  // 文章内容

	LookCount     int `json:"look_count"`     // 文章浏览量
	CommentCount  int `json:"comment_count"`  // 文章评论量
	DiggCount     int `json:"digg_count"`     // 文章点赞量
	CollectsCount int `json:"collects_count"` // 文章收藏量

	// --------用户 一对多 文章--------
	UserID       uint   `json:"user_id"`        // 用户ID
	UserNickName string `json:"user_nick_name"` // 文章作者
	UserAvatar   string `json:"user_avatar"`    // 文章作者头像

	Category string `json:"category"` // 文章分类
	Source   string `json:"source"`   // 文章来源
	Link     string `json:"link"`     // 原文链接

	// --------文章 belongs to 封面--------
	BannerID  uint   `json:"banner_id"`  // 文章封面ID
	BannerUrl string `json:"banner_url"` // 文章封面

	Tags ctype.Array `json:"tags"` // 文章标签
}

func (this ArticleModel) Index() string {
	return "article_index"
}

func (this ArticleModel) Mapping() string {
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
			"abstract": {
				"type":	"text"
			},
			"content": {
				"type":	"text"
			},
			"look_count": {
				"type": "integer"
			},
			"comment_count": {
				"type": "integer"
			},
			"digg_count": {
				"type": "integer"
			},
			"collects_count": {
				"type": "integer"
			},
			"user_id": {
				"type": "integer"
			},
			"user_nick_name": {
				"type":	"text"
			},
			"user_avatar": {
				"type":	"text"
			},
			"category": {
				"type":	"text"
			},
			"source": {
				"type":	"text"
			},
			"link": {
				"type":	"text"
			},
			"banner_id": {
				"type": "integer"
			},
			"banner_url": {
				"type":	"text"
			},
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
		logrus.Error(err.Error())
	}
	return exists
}

// CreateIndex 创建索引
func (this ArticleModel) CreateIndex() error {
	if this.IndexExists() {
		// 索引已经存在，删除索引
		return this.RemoveIndex()
	}
	// 没有索引，创建索引
	createIndex, err := global.ESClient.
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
func (this ArticleModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引")
	indexDelete, err := global.ESClient.
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
