package es_service

import (
	"github.com/olivere/elastic/v7"
	"gvb_server/models"
)

type Option struct {
	models.PageInfo
	Fields   []string
	Tag      string
	Category string
	Query    *elastic.BoolQuery
}

type SortField struct {
	Field     string
	Ascending bool
}

func (o *Option) GetForm() int {
	if o.Page == 0 {
		o.Page = 1
	}
	if o.Limit == 0 {
		o.Limit = 10
	}
	return (o.Page - 1) * o.Limit
}
