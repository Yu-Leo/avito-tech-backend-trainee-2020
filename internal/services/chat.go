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
	chatRepository repositories.ChatRepository
	userRepository repositories.UserRepository
}

func NewChatService(chatRepository repositories.ChatRepository, userRepository repositories.UserRepository) *ChatService {
	return &ChatService{
		chatRepository: chatRepository,
		userRepository: userRepository,
	}
}

func (s ChatService) CreateChat(chat models.CreateChatRequest) (*models.ChatId, error) {
	if utf8.RuneCountInString(chat.Name) > maxLenOfChatName {
		return nil, apperror.TooLongName
	}
	return s.chatRepository.Create(context.Background(), chat)
}

func (s ChatService) GetUserChats(chat models.GetUserChatsRequest) (*[]models.GetUserChatsResponse, error) {
	b, err := s.userRepository.DoesUserIdExist(context.Background(), chat.User)
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, apperror.IDNotFound
	}
	return s.chatRepository.GetUserChats(context.Background(), chat)
}
