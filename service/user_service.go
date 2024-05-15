package service

import (
	"belajar-rest-gorm/model/entity"
	"belajar-rest-gorm/model/web"
)

type UserService interface {
	SaveUser(request web.UserServiceRequest) (map[string]interface{}, error)
	GetUser(userId int) (entity.UserEntity, error)
	GetUseList() ([]entity.UserEntity, error)
	UpdateUser(request web.UserUpdateServiceRequest, pathId int) (map[string]interface{}, error)
	LoginUser(email string, password string) (map[string]interface{}, error)
}
