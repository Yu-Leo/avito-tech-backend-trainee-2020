package repository

import (
	"context"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
)

type MessageRepository interface {
	Create(context.Context, models.CreateMessageDTO) (int, error)
}

type ChatRepository interface {
	Create(context.Context, models.CreateChatDTO) (int, error)
	GetUserChats(context.Context, models.GetUserChatsDTORequest) ([]models.GetUserChatsDTOAnswer, error)
}

type UserRepository interface {
	Create(context.Context, models.CreateUserDTO) (int, error)
}
