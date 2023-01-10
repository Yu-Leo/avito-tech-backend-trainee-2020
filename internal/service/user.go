package service

import (
	"context"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		repository: userRepository,
	}
}

func (s UserService) CreateUser(user models.CreateUserDTO) (int, error) {
	return s.repository.Create(context.Background(), user)
}
