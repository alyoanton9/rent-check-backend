package service

import (
	"github.com/golang-jwt/jwt"
	"math/rand"
)

type AuthService interface {
	GetAuthScheme() string
	CreateToken(login string) (string, error)
}

type authService struct {
	secret []byte
}

func NewAuthService(secret string) AuthService {
	return authService{secret: []byte(secret)}
}

func (s authService) GetAuthScheme() string {
	return "Bearer"
}

func (s authService) CreateToken(login string) (string, error) {
	salt := rand.Int()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"login": login,
			"salt":  salt,
		})

	tokenString, err := token.SignedString(s.secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
