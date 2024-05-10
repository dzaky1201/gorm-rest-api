package service

import (
	"belajar-rest-gorm/model/entity"
	"belajar-rest-gorm/model/web"
)

type UserService interface {
	SaveUser(request web.UserServiceRequest) (map[string]interface{}, error)
	GetUser(userId int) (entity.UserEntity, error)
}
