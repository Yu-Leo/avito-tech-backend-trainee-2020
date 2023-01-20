package services

import (
	"context"
	"unicode/utf8"

	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/apperror"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repositories"
)

const (
	maxLenOfChatName = 80
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
	if utf8.RuneCountInString(chat.Name) > maxLenOfChatName {
		return nil, apperror.TooLongName
	}
	return s.repository.Create(context.Background(), chat)
}

func (s ChatService) GetUserChats(chat models.GetUserChatsRequest) (*[]models.GetUserChatsResponse, error) {
	return s.repository.GetUserChats(context.Background(), chat)
}
