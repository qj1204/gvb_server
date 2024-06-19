package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	core.InitConf()

	fmt.Println(global.Config.BigModel)
}
