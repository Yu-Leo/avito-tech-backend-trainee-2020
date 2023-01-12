package services

import (
	"context"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repositories"
)

type UserService struct {
	repository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		repository: userRepository,
	}
}

func (s UserService) CreateUser(user models.CreateUserDTO) (int, error) {
	return s.repository.Create(context.Background(), user)
}
