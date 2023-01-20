package psql

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/apperror"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repositories"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/postgresql"
)

type chatRepository struct {
	postgresConnection postgresql.Connection
}

func NewPostgresChatRepository(pc postgresql.Connection) repositories.ChatRepository {
	return &chatRepository{
		postgresConnection: pc,
	}
}

func (cr *chatRepository) Create(ctx context.Context, chat models.CreateChatRequest) (chatId *models.ChatId, err error) {
	chatId = &models.ChatId{}

	transaction, err := cr.postgresConnection.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer transaction.Rollback(context.Background())

	chatId.Id, err = createChat(transaction, ctx, chat.Name)
	if err != nil {
		return nil, err
	}

	err = addUsersToChat(transaction, ctx, chatId.Id, chat.Users)
	if err != nil {
		return nil, err
	}

	err = transaction.Commit(context.Background())
	if err != nil {
		return nil, err
	}

	return chatId, nil
}

func createChat(transaction pgx.Tx, ctx context.Context, chatName string) (chatId int, err error) {
	var pgErr *pgconn.PgError
	q := `
INSERT INTO chats (name)
VALUES ($1)
RETURNING chats.id;`
	err = transaction.QueryRow(ctx, q, chatName).Scan(&chatId)
	if err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return 0, apperror.ChatNameAlreadyExists
		}
		return 0, err
	}
	return chatId, nil
}

func addUsersToChat(transaction pgx.Tx, ctx context.Context, chatId int, chatUsers []int) (err error) {
	var pgErr *pgconn.PgError
	q := `
INSERT INTO users_chats (user_id, chat_id)
VALUES ($1, $2);`
	for _, userID := range chatUsers {
		_, err = transaction.Exec(ctx, q, userID, chatId)

		if err != nil {
			if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.ForeignKeyViolation {
				return apperror.IDNotFound
			}
			return err
		}
	}
	return nil
}

func (cr *chatRepository) GetUserChats(ctx context.Context, chat models.GetUserChatsRequest) (*[]models.GetUserChatsResponse, error) {
	transaction, err := cr.postgresConnection.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer transaction.Rollback(context.Background())

	userChats, err := getUserChatsWithoutUsersList(transaction, ctx, chat.User)
	if err != nil {
		return nil, err
	}

	err = addUsersListForChats(transaction, ctx, userChats)
	if err != nil {
		return nil, err
	}

	err = transaction.Commit(context.Background())
	if err != nil {
		return nil, err
	}

	return userChats, nil
}

func getUserChatsWithoutUsersList(transaction pgx.Tx, ctx context.Context, userId int) (*[]models.GetUserChatsResponse, error) {
	answer := make([]models.GetUserChatsResponse, 0)
	q := `
SELECT chats.id, chats.name, chats.created_at
FROM users_chats
JOIN chats on users_chats.chat_id = chats.id
WHERE users_chats.user_id = $1
ORDER BY (SELECT (MAX(created_at))
          FROM messages
          WHERE chat_id = users_chats.chat_id
          GROUP BY chat_id) DESC;
`
	rows, err := transaction.Query(ctx, q, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		userChat := models.GetUserChatsResponse{}
		err = rows.Scan(&userChat.Id, &userChat.Name, &userChat.CreatedAt)
		if err != nil {
			return nil, err
		}
		answer = append(answer, userChat)
	}
	rows.Close()
	return &answer, nil
}

func addUsersListForChats(transaction pgx.Tx, ctx context.Context, chats *[]models.GetUserChatsResponse) error {
	for i := range *chats {
		q := `
SELECT users_chats.user_id
FROM users_chats
WHERE users_chats.chat_id = $1;
`
		rows, err := transaction.Query(ctx, q, (*chats)[i].Id)
		if err != nil {
			return err
		}
		for rows.Next() {
			var userId int
			err = rows.Scan(&userId)
			if err != nil {
				return err
			}
			(*chats)[i].Users = append((*chats)[i].Users, userId)
		}
		rows.Close()
	}
	return nil
}
