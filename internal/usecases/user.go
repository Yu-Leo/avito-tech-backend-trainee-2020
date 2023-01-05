package usecases

import (
	"context"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/entities"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repositories"
)

type UserUseCase struct {
	repository repositories.UserRepository
}

func NewUserUserCase(userRepository repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		repository: userRepository,
	}
}

func (u UserUseCase) CreateUser(user entities.UserDTO) (int, error) {
	return u.repository.CreateUser(context.Background(), user)
}
