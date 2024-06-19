package common_service

import (
	"fmt"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.Page
	Debug   bool
	Likes   []string // 模糊匹配的字段
	Where   *gorm.DB // 额外的查询
	Preload []string // 预加载的字段列表
}

func CommonList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	// 如果是debug模式，就打印sql（MysqlLog始终为显示所有sql）
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	// 默认按创建时间降序排列
	if option.Sort == "" {
		option.Sort = "created_at desc"
	}

	// 这样可以查询model里的字段
	query := DB.Where(model)

	// 如果有高级查询就加上
	if option.Where != nil {
		query.Where(option.Where)
	}

	// 模糊查询
	if option.Key != "" {
		for index, cloumn := range option.Likes {
			if index == 0 { // 第一个like
				query = query.Where(fmt.Sprintf("%s like ?", cloumn), "%"+option.Key+"%")
			} else { // 后面的like
				query = query.Or(fmt.Sprintf("%s like ?", cloumn), "%"+option.Key+"%")
			}
		}
	}
	q1 := query

	// 查列表，获取总数
	count = query.Find(&list).RowsAffected

	// 预加载
	for _, preload := range option.Preload {
		q1 = q1.Preload(preload)
	}

	offset := option.Limit * (option.PageNum - 1) // 由前端传过来，PageNum肯定不为0
	offset = max(offset, 0)
	if option.Limit == 0 {
		option.Limit = -1
	}

	// 当limit为-1时，表示不分页，下面的limit不生效
	err = q1.Offset(offset).Limit(option.Limit).Order(option.Sort).Find(&list).Error
	return
}
