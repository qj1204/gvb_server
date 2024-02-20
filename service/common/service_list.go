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
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按创建时间倒序排列
	}

	query := DB.Where(model)
	count = query.Select("id").Find(&list).RowsAffected
	// 这里的query会受上面的查询的影响，需要手动复位
	query = DB.Where(model)
	offset := option.Limit * (option.PageNum - 1)
	offset = max(offset, 0)
	if option.Limit == 0 {
		option.Limit = -1
	}
	err = query.Offset(offset).Limit(option.Limit).Order(option.Sort).Find(&list).Error
	return
}
