package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/service/es_service"
)

func main() {
	core.InitConf()
	core.InitLogger()
	global.ESClient = core.EsConnect()
	es_service.DeleteFullTextByArticleID("MI4aeYYB6uoytGZAtrHU")

}
