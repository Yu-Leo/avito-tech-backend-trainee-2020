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

type chatRepository struct {
	postgresConnection postgresql.Connection
}

func NewPostgresChatRepository(pc postgresql.Connection) repositories.ChatRepository {
	return &chatRepository{
		postgresConnection: pc,
	}
}

func (cr *chatRepository) Create(ctx context.Context, chat models.CreateChatDTO) (chatId models.ChatId, err error) {
	var pgErr *pgconn.PgError

	transaction, err := cr.postgresConnection.Begin(context.Background())
	if err != nil {
		return models.ChatId{}, err
	}
	defer transaction.Rollback(context.Background())

	q1 := `
INSERT INTO chats (name)
VALUES ($1)
RETURNING chats.id;
				`
	err = transaction.QueryRow(ctx, q1, chat.Name).Scan(&chatId.Id)
	if err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return models.ChatId{}, apperror.ChatNameAlreadyExists
		}
		return models.ChatId{}, err
	}
	q2 := `
INSERT INTO users_chats (user_id, chat_id)
VALUES ($1, $2);
				`
	for _, userID := range chat.Users {
		_, err = transaction.Exec(ctx, q2, userID, chatId.Id)

		if err != nil {
			if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.ForeignKeyViolation {
				return models.ChatId{}, apperror.IDNotFound
			}
			return models.ChatId{}, err
		}

	}
	err = transaction.Commit(context.Background())
	if err != nil {
		return models.ChatId{}, err
	}

	return chatId, nil
}

func (cr *chatRepository) GetUserChats(ctx context.Context, chat models.GetUserChatsDTORequest) (*[]models.GetUserChatsDTOAnswer, error) {
	transaction, err := cr.postgresConnection.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer transaction.Rollback(context.Background())

	answer := make([]models.GetUserChatsDTOAnswer, 0)
	q1 := `
SELECT chats.id, chats.name, chats.created_at
FROM users_chats
JOIN chats on users_chats.chat_id = chats.id
WHERE users_chats.user_id = $1
ORDER BY (SELECT (MAX(created_at))
          FROM messages
          WHERE chat_id = users_chats.chat_id
          GROUP BY chat_id) DESC;
`
	rows, err := transaction.Query(ctx, q1, chat.User)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		userChat := models.GetUserChatsDTOAnswer{}
		err = rows.Scan(&userChat.Id, &userChat.Name, &userChat.CreatedAt)
		if err != nil {
			return nil, err
		}
		answer = append(answer, userChat)
	}
	rows.Close()

	for i := range answer {
		q2 := `
SELECT users_chats.user_id
FROM users_chats
WHERE users_chats.chat_id = $1;
`
		chatUsers, err := transaction.Query(ctx, q2, answer[i].Id)
		if err != nil {
			return nil, err
		}
		for chatUsers.Next() {
			var userId int
			err = chatUsers.Scan(&userId)
			if err != nil {
				return nil, err
			}
			answer[i].Users = append(answer[i].Users, userId)
		}
	}
	rows.Close()

	err = transaction.Commit(context.Background())
	if err != nil {
		return nil, err
	}

	return &answer, nil
}
