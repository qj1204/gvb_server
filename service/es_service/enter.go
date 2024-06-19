package es_service

import (
	"github.com/olivere/elastic/v7"
	"gvb_server/models"
)

type Option struct {
	models.Page
	Fields   []string
	Tag      string
	Category string
	Query    *elastic.BoolQuery
}

type SortField struct {
	Field     string
	Ascending bool
}

func (o *Option) GetFrom() int {
	if o.Limit == 0 { // 默认每页10条
		o.Limit = 10
	}
	if o.PageNum == 0 {
		o.PageNum = 1
	}
	return (o.PageNum - 1) * o.Limit
}
