package repository

import "belajar-rest-gorm/model/domain"

type UserRepository interface {
	SaveUser(user domain.User) (domain.User, error)
	GetUser(Id int) (domain.User, error)
	GetUsers() ([]domain.User, error)
	UpdateUser(user domain.User) (domain.User, error)
	FindUserByEmail(email string) (*domain.User, error)
}
