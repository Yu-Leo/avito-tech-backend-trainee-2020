package services

import (
	"context"
	"unicode/utf8"

	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/apperror"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repositories"
)

const (
	maxLenOfUserName = 80
)

type UserService struct {
	repository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		repository: userRepository,
	}
}

func (s UserService) CreateUser(requestData models.CreateUserRequest) (*models.UserId, error) {
	if utf8.RuneCountInString(requestData.Username) > maxLenOfUserName {
		return nil, apperror.TooLongName
	}
	return s.repository.Create(context.Background(), requestData)
}
