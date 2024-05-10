package entity

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
