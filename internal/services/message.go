package services

import (
	"context"

	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repositories"
)

type MessageService struct {
	repository repositories.MessageRepository
}

func NewMessageService(messageRepository repositories.MessageRepository) *MessageService {
	return &MessageService{
		repository: messageRepository,
	}
}

func (s MessageService) CreateMessage(message models.CreateMessageRequest) (*models.MessageId, error) {
	return s.repository.Create(context.Background(), message)
}

func (s MessageService) GetChatMessages(chat models.GetChatMessagesRequest) (*[]models.Message, error) {
	return s.repository.GetChatMessages(context.Background(), chat)
}
