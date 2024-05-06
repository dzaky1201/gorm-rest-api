package service

import (
	"belajar-rest-gorm/model/web"
)

type UserService interface {
	SaveUser(request web.UserServiceRequest)(map[string]interface{},error)
}
