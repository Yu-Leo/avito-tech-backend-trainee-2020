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

func (s ChatService) CreateChat(requestData models.CreateChatRequest) (*models.ChatId, error) {
	if utf8.RuneCountInString(requestData.Name) > maxLenOfChatName {
		return nil, apperror.TooLongName
	}
	return s.chatRepository.Create(context.Background(), requestData)
}

func (s ChatService) GetUserChats(requestData models.GetUserChatsRequest) (*[]models.GetUserChatsResponse, error) {
	b, err := s.userRepository.DoesUserIdExist(context.Background(), requestData.User)
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, apperror.IDNotFound
	}
	return s.chatRepository.GetUserChats(context.Background(), requestData)
}
