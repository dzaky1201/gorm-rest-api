package service

import (
	"belajar-rest-gorm/repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	repository repository.UserRepository
	Validate   *validator.Validate
}

func NewUserService(repository repository.UserRepository, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{repository: repository, Validate: validate}
}
