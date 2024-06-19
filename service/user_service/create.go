package user_service

import (
	"errors"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils"
	"gvb_server/utils/pwd"
)

const Avatar = "/static/avatar/default.jpg"

func (this *UserService) CreateUser(nickName, userName, password string, role ctype.Role, email, ip string) error {
	// 判断用户名是否存在
	var userModel models.UserModel
	count := global.DB.Take(&userModel, "user_name = ?", userName).RowsAffected
	if count > 0 {
		// 用户名已存在
		return errors.New("用户名已存在")
	}
	// 对密码进行加密
	hashPwd := pwd.HashPwd(password)

	// 写入数据库
	err := global.DB.Create(&models.UserModel{
		UserName:   userName,
		NickName:   nickName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     Avatar,
		IP:         ip,
		Addr:       utils.GetAddr(ip),
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
