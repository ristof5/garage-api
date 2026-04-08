package services

import (
	"errors"
	"garage-api/helpers"
	"garage-api/repositories"
)

type AuthService struct {
	Repo *repositories.UserRepository
}

func (s *AuthService) Login(username, password string) (string, error) {

	user, err := s.Repo.GetByUsername(username)

	if err != nil {
		return "", errors.New("user not found")
	}

	if !helpers.CheckPassword(password, user.Password) {
		return "", errors.New("invalid password")
	}

	token, err := helpers.GenerateToken(user.ID, user.Role)

	if err != nil {
		return "", err
	}

	return token, nil
}
