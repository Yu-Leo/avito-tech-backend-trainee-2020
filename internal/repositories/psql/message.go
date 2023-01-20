package psql

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/apperror"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repositories"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/postgresql"
)

type messageRepository struct {
	postgresConnection postgresql.Connection
}

func NewPostgresMessageRepository(pc postgresql.Connection) repositories.MessageRepository {
	return &messageRepository{
		postgresConnection: pc,
	}
}

func (mr *messageRepository) Create(ctx context.Context, message models.CreateMessageRequest) (messageId *models.MessageId, err error) {
	messageId = &models.MessageId{}

	userInChat, err := isUserInChat(mr.postgresConnection, ctx, message)
	if err != nil {
		return nil, err
	}
	if !userInChat {
		return nil, apperror.UserIsNotInChat
	}

	q2 := `
INSERT INTO messages (user_id, chat_id, message_text)
VALUES ($1, $2, $3)
RETURNING messages.id;`
	err = mr.postgresConnection.QueryRow(ctx, q2, message.UserId, message.ChatId, message.Text).Scan(&(*messageId).Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.ForeignKeyViolation {
			return nil, apperror.IDNotFound
		}
		return nil, err
	}
	return messageId, nil
}

func isUserInChat(conn postgresql.Connection, ctx context.Context, message models.CreateMessageRequest) (bool, error) {
	q1 := `
SELECT id
FROM users_chats
WHERE user_id = $1 AND chat_id = $2
`
	rows, err := conn.Query(ctx, q1, message.UserId, message.ChatId)

	if err != nil {
		return false, err
	}
	if !rows.Next() {
		return false, nil
	}
	rows.Close()
	return true, nil
}

func (mr *messageRepository) GetChatMessages(ctx context.Context,
	chat models.GetChatMessagesRequest) (*[]models.Message, error) {
	answer := make([]models.Message, 0)

	q := `
SELECT id, user_id, chat_id, message_text, created_at
FROM messages
WHERE chat_id = $1

`
	rows, err := mr.postgresConnection.Query(ctx, q, chat.ChatId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		message := models.Message{}
		err = rows.Scan(&message.Id, &message.UserId, &message.ChatId, &message.Text, &message.CreatedAt)

		if err != nil {
			return nil, err
		}
		answer = append(answer, message)
	}
	rows.Close()

	return &answer, nil
}
