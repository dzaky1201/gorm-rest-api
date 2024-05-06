package service

import "belajar-rest-gorm/model/web"

type UserService interface {
	SaveUser(req web.UserServiceRequest) (web.WebResponse, error)
}
