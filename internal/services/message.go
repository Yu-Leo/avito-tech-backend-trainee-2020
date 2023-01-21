package services

import (
	"context"

	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/apperror"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repositories"
)

type MessageService struct {
	messageRepository repositories.MessageRepository
	chatRepository    repositories.ChatRepository
}

func NewMessageService(messageRepository repositories.MessageRepository, chatRepository repositories.ChatRepository) *MessageService {
	return &MessageService{
		messageRepository: messageRepository,
		chatRepository: chatRepository,
	}
}

func (s MessageService) CreateMessage(message models.CreateMessageRequest) (*models.MessageId, error) {
	return s.messageRepository.Create(context.Background(), message)
}

func (s MessageService) GetChatMessages(chat models.GetChatMessagesRequest) (*[]models.Message, error) {
	b, err := s.chatRepository.DoesChatIdExist(context.Background(), chat.ChatId)
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, apperror.IDNotFound
	}
	return s.messageRepository.GetChatMessages(context.Background(), chat)
}
