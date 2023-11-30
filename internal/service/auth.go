package service

import "fmt"

type AuthService interface {
	AuthenticateUser(authToken string, refreshToken string) (string, error)
	GetAuthScheme() string
}

type authService struct{}

func NewAuthService() AuthService {
	return authService{}
}

func (s authService) GetAuthScheme() string {
	return "OAuth"
}

func (s authService) AuthenticateUser(authToken string, refreshToken string) (string, error) {
	// TODO: implementation to use Google Identity and others
	switch authToken {
	case "test":
		return "0", nil
	case "test1":
		return "1", nil
	case "test2":
		return "2", nil
	case "test3":
		return "3", nil
	case "test4":
		return "4", nil
	case "test5":
		return "5", nil
	default:
		return "", fmt.Errorf("invalid auth token")
	}
}
