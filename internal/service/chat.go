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

func (s ChatService) CreateChat(chat models.CreateChatDTO) (int, error) {
	return s.repository.Create(context.Background(), chat)
}

func (s ChatService) GetUserChats(chat models.GetUserChatsDTORequest) ([]models.GetUserChatsDTOAnswer, error) {
	return s.repository.GetUserChats(context.Background(), chat)
}
