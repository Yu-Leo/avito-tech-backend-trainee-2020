package service

import (
	"context"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repository"
)

type MessageService struct {
	repository repository.MessageRepository
}

func NewMessageService(messageRepository repository.MessageRepository) *MessageService {
	return &MessageService{
		repository: messageRepository,
	}
}

func (s MessageService) CreateMessage(message models.MessageDTO) (int, error) {
	return s.repository.Create(context.Background(), message)
}
