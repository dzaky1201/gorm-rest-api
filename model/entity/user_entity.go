package entity

import "belajar-rest-gorm/model/domain"

type UserEntity struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

func ToUserEntity(user_id int, name string, email string) UserEntity {
	return UserEntity{
		UserID: user_id,
		Name:   name,
		Email:  email,
	}
}

func ToUserListEntity(users []domain.User) []UserEntity {
	usersData := []UserEntity{}

	for _, user := range users {
		usersData = append(usersData, ToUserEntity(user.UserID, user.Name, user.Email))
	}

	return usersData
}
