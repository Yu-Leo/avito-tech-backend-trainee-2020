package service

import (
	"context"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repository"
)

type ChatService struct {
	repository repository.ChatRepository
}

func NewChatService(chatRepository repository.ChatRepository) *ChatService {
	return &ChatService{
		repository: chatRepository,
	}
}

func (s ChatService) CreateChat(chat models.ChatDTO) (int, error) {
	return s.repository.Create(context.Background(), chat)
}
