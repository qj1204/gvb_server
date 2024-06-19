package user_service

import (
	"gvb_server/service/redis_service"
)

func (this *UserService) BindEmail(token, email, code string) error {
	return redis_service.BindEmail(token, email, code)
}

func (this *UserService) GetEmailAndCodeByToken(token string) (string, string, error) {
	return redis_service.GetEmailAndCodeByToken(token)
}

func (this *UserService) DelEmailAndCode(token string) error {
	return redis_service.DelEmailAndCode(token)
}
