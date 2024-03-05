package pwd

import (
	"golang.org/x/crypto/bcrypt"
	"gvb_server/global"
)

// HashPwd 加密密码
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		global.Log.Error(err.Error())
	}
	return string(hash)
}

// CheckPwd 验证加密后的密码
func CheckPwd(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		global.Log.Error(err.Error())
		return false
	}
	return true
}
