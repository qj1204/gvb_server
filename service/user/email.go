package user

import (
	"gvb_server/service/redis"
)

func (this *UserService) BindEmail(token, email, code string) error {
	return redis.BindEmail(token, email, code)
}

func (this *UserService) GetEmailAndCodeByToken(token string) (string, string, error) {
	return redis.GetEmailAndCodeByToken(token)
}

func (this *UserService) DelEmailAndCode(token string) error {
	return redis.DelEmailAndCode(token)
}
