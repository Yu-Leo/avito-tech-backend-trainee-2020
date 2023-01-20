package services

import (
	"context"

	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repositories"
)

type ChatService struct {
	repository repositories.ChatRepository
}

func NewChatService(chatRepository repositories.ChatRepository) *ChatService {
	return &ChatService{
		repository: chatRepository,
	}
}

func (s ChatService) CreateChat(chat models.CreateChatRequest) (*models.ChatId, error) {
	return s.repository.Create(context.Background(), chat)
}

func (s ChatService) GetUserChats(chat models.GetUserChatsRequest) (*[]models.GetUserChatsResponse, error) {
	return s.repository.GetUserChats(context.Background(), chat)
}
