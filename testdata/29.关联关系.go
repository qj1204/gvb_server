package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
)

func main() {
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()

	var model []models.BigModelTagModel

	global.DB.Preload("Roles").Find(&model, []uint{5})
	//fmt.Println(model.Roles)
	err := global.DB.Debug().Model(&model).Association("Roles").Delete(model)
	fmt.Println(err)

}
