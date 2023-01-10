package repository

import (
	"context"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
)

type MessageRepository interface {
	Create(context.Context, models.MessageDTO) (int, error)
}

type ChatRepository interface {
	Create(context.Context, models.ChatDTO) (int, error)
}

type UserRepository interface {
	Create(context.Context, models.UserDTO) (int, error)
}
