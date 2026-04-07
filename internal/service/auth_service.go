package service

import (
	"errors"

	"github.com/Aditya7880900936/credes-backend/internal/models"
	"github.com/Aditya7880900936/credes-backend/internal/repository"
	"github.com/Aditya7880900936/credes-backend/internal/utils"
)

func RegisterUser(email, password, fullName string) (*models.User, error) {
	hash, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    email,
		Password: hash,
		FullName: fullName,
		Role:     models.RoleUser,
	}

	err = repository.CreateUser(user)
	return user, err
}

func LoginUser(email, password, secret string) (string, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !user.IsActive {
		return "", errors.New("user is inactive")
	}

	err = utils.CheckPassword(password, user.Password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateToken(user.ID, string(user.Role), secret)
}
