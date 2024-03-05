package desens

import "strings"

// DesensitizationEmail 邮箱脱敏
func DesensitizationEmail(email string) string {
	eList := strings.Split(email, "@")
	if len(eList) != 2 {
		return ""
	}
	return eList[0][:1] + "****@" + eList[1]
}

// DesensitizationTel 电话脱敏
func DesensitizationTel(tel string) string {
	if len(tel) < 11 {
		return tel
	}
	return tel[:3] + "****" + tel[7:]
}
