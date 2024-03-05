package common

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.Page
	Debug bool
}

func CommonList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	// 如果是debug模式，就打印sql（MysqlLog始终为显示所有sql）
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按创建时间降序排列
	}

	query := DB.Where(model) // 这样可以查询model里的字段
	count = query.Select("id").Find(&list).RowsAffected
	// 这里的query会受上面的查询的影响，需要手动复位
	query = DB.Where(model)
	offset := option.Limit * (option.PageNum - 1) // 由前端传过来，PageNum肯定不为0
	offset = max(offset, 0)
	if option.Limit == 0 {
		option.Limit = -1
	}
	// 当limit为-1时，表示不分页，下面的limit不生效
	err = query.Offset(offset).Limit(option.Limit).Order(option.Sort).Find(&list).Error
	return
}
