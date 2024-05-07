package service

import (
	"belajar-rest-gorm/helper"
	"belajar-rest-gorm/model/domain"
	"belajar-rest-gorm/model/web"
	"belajar-rest-gorm/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type ResponseToJson map[string]interface{}

type UserServiceImpl struct {
	repository repository.UserRepository
	validate   *validator.Validate
}

func NewUserService(repository repository.UserRepository, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		repository: repository,
		validate:   validate,
	}
}

func (service *UserServiceImpl) SaveUser(request web.UserServiceRequest) (map[string]interface{}, error) {
	if err := service.validate.Struct(request); err != nil {
		return nil, err
	}

	passHash, errHash := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)

	if errHash != nil {
		return nil, errHash
	}

	userReq := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(passHash),
	}

	saveUser, errSaveUser := service.repository.SaveUser(userReq)

	if errSaveUser != nil {
		return nil, errSaveUser
	}

	return helper.ResponseToJson{"name": saveUser.Name, "email": saveUser.Email}, nil

}
