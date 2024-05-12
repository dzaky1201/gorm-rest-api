package repository

import (
	"belajar-rest-gorm/model/domain"
	"errors"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (repo *UserRepositoryImpl) SaveUser(user domain.User) (domain.User, error) {

	err := repo.db.Create(&user).Error

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repo *UserRepositoryImpl) GetUser(Id int) (domain.User, error) {
	var userData domain.User

	err := repo.db.First(&userData, "user_id = ?", Id).Error

	if err != nil {
		return domain.User{}, errors.New("user tidak ditemukan")
	}

	return userData, nil
}

func (repo *UserRepositoryImpl) GetUsers() ([]domain.User, error) {
	var users []domain.User

	err := repo.db.Find(&users).Error

	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

func (repo *UserRepositoryImpl)UpdateUser(user domain.User)(domain.User, error){
	err := repo.db.Model(domain.User{}).Where("user_id = ?", user.UserID).Updates(user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
