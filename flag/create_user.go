package flag

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models/common/ctype"
	"gvb_server/service/user"
)

func CreateUser(permission string) {
	// 用户名 昵称 密码 确认密码 邮箱
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

	// 判断密码是否一致
	if password != rePassword {
		global.Log.Error("两次密码不一致，请重新输入")
		return
	}
	// 角色
	role := ctype.PermissionUser
	if permission == "admin" {
		role = ctype.PermissionAdmin
	}
	s := &user.UserService{}
	err := s.CreateUser(nickName, userName, password, role, email, "127.0.0.1")
	if err != nil {
		global.Log.Error("创建用户失败", err)
		return
	}
	global.Log.Info("创建用户成功")
}
