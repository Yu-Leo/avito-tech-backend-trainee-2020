package psql

import (
	"context"
	"errors"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/apperror"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repositories"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/postgresql"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"time"
)

type chatRepository struct {
	postgresConnection postgresql.Connection
}

func NewPostgresChatRepository(pc postgresql.Connection) repositories.ChatRepository {
	return &chatRepository{
		postgresConnection: pc,
	}
}

func (cr *chatRepository) Create(ctx context.Context, chat models.CreateChatDTO) (chatID int, err error) {
	transaction, err := cr.postgresConnection.Begin(context.Background())
	if err != nil {
		return 0, err
	}
	defer transaction.Rollback(context.Background())

	q1 := `
INSERT INTO chats (name)
VALUES ($1)
RETURNING chats.id;
				`
	err = transaction.QueryRow(ctx, q1, chat.Name).Scan(&chatID)
	if err != nil {
		return 0, err
	}
	q2 := `
INSERT INTO users_chats (user_id, chat_id)
VALUES ($1, $2);
				`
	for _, userID := range chat.Users {
		_, err = transaction.Exec(ctx, q2, userID, chatID)

		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.ForeignKeyViolation {
				return 0, apperror.IDNotFound
			}
			return 0, err
		}

	}
	err = transaction.Commit(context.Background())
	if err != nil {
		return 0, err
	}

	return chatID, nil
}

func (cr *chatRepository) GetUserChats(context.Context, models.GetUserChatsDTORequest) ([]models.GetUserChatsDTOAnswer, error) {
	answer := make([]models.GetUserChatsDTOAnswer, 1)

	users := make([]int, 3)
	users[0] = 1
	users[1] = 2
	users[2] = 3

	answer[0] = models.GetUserChatsDTOAnswer{
		Id:        0,
		Name:      "Chat 0",
		Users:     users,
		CreatedAt: time.Time{},
	}
	return answer, nil
}
