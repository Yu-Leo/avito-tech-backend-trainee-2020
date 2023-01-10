package psql

import (
	"context"
	"errors"
	"fmt"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/apperror"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repository"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/postgresql"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type messageRepository struct {
	postgresConnection postgresql.Connection
}

func NewPostgresMessageRepository(pc postgresql.Connection) repository.MessageRepository {
	return &messageRepository{
		postgresConnection: pc,
	}
}

func (mr *messageRepository) Create(ctx context.Context, chat models.MessageDTO) (messageID int, err error) {
	q := `
INSERT INTO messages (user_id, chat_id, message_text)
VALUES ($1, $2, $3)
RETURNING messages.id;`
	err = mr.postgresConnection.QueryRow(ctx, q, chat.UserId, chat.ChatId, chat.Text).Scan(&messageID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.ForeignKeyViolation {
			fmt.Println(pgErr.Error())
			fmt.Println(pgErr.Message)
			return 0, apperror.IDNotFound
		}
		return 0, err
	}
	return messageID, nil
}
