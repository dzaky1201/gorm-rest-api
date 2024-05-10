package web

type UserUpdateServiceRequest struct {
	Name  string `json:"name"`
	Email string `validate:"email" json:"email"`
}