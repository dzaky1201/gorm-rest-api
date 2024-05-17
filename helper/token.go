package helper

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type TokenUseCase interface {
	GenerateAccessToken(claims JwtCustomClaims) (string, error)
}

type TokenUseCaseImpl struct{}

type JwtCustomClaims struct {
	ID    string `json:"user_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func NewTokenUseCase() *TokenUseCaseImpl {
	return &TokenUseCaseImpl{}
}

func (t *TokenUseCaseImpl) GenerateAccessToken(claims JwtCustomClaims) (string, error) {

	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := plainToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return encodedToken, nil
}
