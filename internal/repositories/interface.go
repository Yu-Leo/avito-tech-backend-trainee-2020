package repositories

import (
	"context"

	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
)

type MessageRepository interface {
	Create(context.Context, models.CreateMessageRequest) (*models.MessageId, error)
	GetChatMessages(context.Context, models.GetChatMessagesRequest) (*[]models.Message, error)
}

type ChatRepository interface {
	Create(context.Context, models.CreateChatRequest) (*models.ChatId, error)
	GetUserChats(context.Context, models.GetUserChatsRequest) (*[]models.GetUserChatsResponse, error)
}

type UserRepository interface {
	Create(context.Context, models.CreateUserRequest) (*models.UserId, error)
}
