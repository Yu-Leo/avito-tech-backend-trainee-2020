package usecases

import (
	"context"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/entities"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repositories"
)

type ChatUseCase struct {
	repository repositories.ChatRepository
}

func NewChatUserCase(chatRepository repositories.ChatRepository) *ChatUseCase {
	return &ChatUseCase{
		repository: chatRepository,
	}
}

func (u ChatUseCase) CreateChat(chat entities.ChatDTO) (int, error) {
	return u.repository.Create(context.Background(), chat)
}
