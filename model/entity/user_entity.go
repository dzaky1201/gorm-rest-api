package entity

import (
	"belajar-rest-gorm/model/domain"
)

type UserEntity struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Address  interface{} `json:"address"`
}


func ToUserEntity(user domain.User) UserEntity {

	if user.Address != nil {
		address := AddressEntity{
			ID: user.Address.AddressID,
			City: user.Address.City,
			Province: user.Address.Province,
			PostalCode: user.Address.PostalCode,
		}

		return UserEntity{
			UserID: user.UserID,
			Name:   user.Name,
			Email:  user.Email,
			Address: address,
		}
	}


	return UserEntity{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
		Address: "alamat belum didaftarkan",
	}

		
	
}

func ToUserListEntity(users []domain.User)[]UserEntity {
	userData := []UserEntity{}

	for _, user := range users {
		userData = append(userData, ToUserEntity(user))
	}

	return userData
}
