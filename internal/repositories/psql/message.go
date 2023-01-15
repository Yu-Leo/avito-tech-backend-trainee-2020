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
)

type messageRepository struct {
	postgresConnection postgresql.Connection
}

func NewPostgresMessageRepository(pc postgresql.Connection) repositories.MessageRepository {
	return &messageRepository{
		postgresConnection: pc,
	}
}

func (mr *messageRepository) Create(ctx context.Context, chat models.CreateMessageDTO) (messageId models.MessageId, err error) {
	q1 := `
SELECT id
FROM users_chats
WHERE user_id = $1 AND chat_id = $2
`
	rows, err := mr.postgresConnection.Query(ctx, q1, chat.UserId, chat.ChatId)

	if err != nil {
		return models.MessageId{}, err
	}
	if !rows.Next() {
		return models.MessageId{}, apperror.UserIsNotInChat
	}
	rows.Close()

	q2 := `
INSERT INTO messages (user_id, chat_id, message_text)
VALUES ($1, $2, $3)
RETURNING messages.id;`
	err = mr.postgresConnection.QueryRow(ctx, q2, chat.UserId, chat.ChatId, chat.Text).Scan(&messageId.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.ForeignKeyViolation {
			return models.MessageId{}, apperror.IDNotFound
		}
		return models.MessageId{}, err
	}
	return messageId, nil
}

func (mr *messageRepository) GetChatMessages(ctx context.Context,
	chat models.GetChatMessagesDRORequest) (*[]models.Message, error) {
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

	return &answer, nil
}
