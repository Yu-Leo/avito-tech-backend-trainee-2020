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

func (s MessageService) CreateMessage(message models.CreateMessageDTO) (int, error) {
	return s.repository.Create(context.Background(), message)
}
