package service

import (
	"gvb_server/service/image"
	"gvb_server/service/user"
)

type ServiceGroup struct {
	ImageService image.ImageService
	UserService  user.UserService
}

var ServiceGroupApp = new(ServiceGroup)
