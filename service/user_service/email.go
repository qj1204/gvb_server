package user_service

import (
	"gvb_server/service/redis_service"
)

func (UserService) BindEmail(token, email, code string) error {
	return redis_service.BindEmail(token, email, code)
}

func (UserService) GetEmailAndCodeByToken(token string) (string, string, error) {
	return redis_service.GetEmailAndCodeByToken(token)
}

func (UserService) DelEmailAndCode(token string) error {
	return redis_service.DelEmailAndCode(token)
}
