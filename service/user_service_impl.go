package service

import (
	"belajar-rest-gorm/helper"
	"belajar-rest-gorm/model/domain"
	"belajar-rest-gorm/model/entity"
	"belajar-rest-gorm/model/web"
	"belajar-rest-gorm/repository"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type ResponseToJson map[string]interface{}

type UserServiceImpl struct {
	repository   repository.UserRepository
	tokenUseCase helper.TokenUseCase
}

func NewUserService(repository repository.UserRepository, token helper.TokenUseCase) *UserServiceImpl {
	return &UserServiceImpl{
		repository:   repository,
		tokenUseCase: token,
	}
}

func (service *UserServiceImpl) SaveUser(request web.UserServiceRequest) (map[string]interface{}, error) {

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

func (service *UserServiceImpl) GetUser(userId int) (entity.UserEntity, error) {
	getUser, errGetUser := service.repository.GetUser(userId)

	if errGetUser != nil {
		return entity.UserEntity{}, errGetUser
	}

	return entity.ToUserEntity(getUser.UserID, getUser.Name, getUser.Email), nil
}

func (service *UserServiceImpl) GetUseList() ([]entity.UserEntity, error) {
	getUserList, errGetUserList := service.repository.GetUsers()

	if errGetUserList != nil {
		return []entity.UserEntity{}, errGetUserList
	}

	return entity.ToUserListEntity(getUserList), nil
}

func (service *UserServiceImpl) UpdateUser(request web.UserUpdateServiceRequest, pathId int) (map[string]interface{}, error) {
	getUserById, err := service.repository.GetUser(pathId)
	if err != nil {
		return nil, err
	}

	if request.Name == "" {
		request.Name = getUserById.Name
	}

	if request.Email == "" {
		request.Email = getUserById.Email
	}

	userRequest := domain.User{
		UserID:   pathId,
		Name:     request.Name,
		Email:    request.Email,
		Password: getUserById.Password,
	}

	updateUser, errUpdate := service.repository.UpdateUser(userRequest)

	if errUpdate != nil {
		return nil, errUpdate
	}

	return helper.ResponseToJson{"name": updateUser.Name, "email": updateUser.Email}, nil
}

func (service *UserServiceImpl) LoginUser(email string, password string) (map[string]interface{}, error) {
	user, err := service.repository.FindUserByEmail(email)

	if err != nil {
		return nil, errors.New("Email tidak ditemukan")
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errPass != nil {
		return nil, errors.New("Password Salah")
	}

	expiredTime := time.Now().Local().Add(1 * time.Hour)

	claims := helper.JwtCustomClaims{
		ID:    strconv.Itoa(user.UserID),
		Name:  user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "rest-gorm",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, errToken := service.tokenUseCase.GenerateAccessToken(claims)

	if errToken != nil {
		return nil, errors.New("ada kesalahan generate token")
	}

	return helper.ResponseToJson{"token": token}, nil
}
