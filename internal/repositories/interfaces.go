package repositories

import (
	"context"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/entities"
)

type UserRepository interface {
	Create(context.Context, entities.UserDTO) (int, error)
}

type ChatRepository interface {
	Create(context.Context, entities.ChatDTO) (int, error)
}
