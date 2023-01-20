package repositories

import (
	"context"

	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
)

type MessageRepository interface {
	Create(context.Context, models.CreateMessageDTO) (*models.MessageId, error)
	GetChatMessages(context.Context, models.GetChatMessagesDRORequest) (*[]models.Message, error)
}

type ChatRepository interface {
	Create(context.Context, models.CreateChatDTO) (*models.ChatId, error)
	GetUserChats(context.Context, models.GetUserChatsDTORequest) (*[]models.GetUserChatsDTOResponse, error)
}

type UserRepository interface {
	Create(context.Context, models.CreateUserDTO) (*models.UserId, error)
}
