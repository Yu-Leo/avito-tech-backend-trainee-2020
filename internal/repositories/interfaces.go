package repositories

import (
	"context"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/entities"
)

type UserRepository interface {
	CreateUser(context.Context, entities.UserDTO) (int, error)
}
