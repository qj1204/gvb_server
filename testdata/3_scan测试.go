package main

import "fmt"

func main() {
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)
	fmt.Printf("请输入用户名：")
	fmt.Scanln(&userName)
	fmt.Printf("请输入昵称：")
	fmt.Scanln(&nickName)
	fmt.Printf("请输入密码：")
	fmt.Scanln(&password)
	fmt.Printf("请确认密码：")
	fmt.Scanln(&rePassword)
	fmt.Printf("请输入邮箱：")
	fmt.Scanln(&email)

	fmt.Println(userName, nickName, password, rePassword, email)
}
