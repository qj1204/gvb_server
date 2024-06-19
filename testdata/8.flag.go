package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string

	flag.StringVar(&user, "u", "", "创建用户")
	flag.Parse()

	fmt.Println(user)
}
