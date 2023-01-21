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

type userRepository struct {
	postgresConnection postgresql.Connection
}

func NewPostgresUserRepository(pc postgresql.Connection) repositories.UserRepository {
	return &userRepository{
		postgresConnection: pc,
	}
}

func (ur *userRepository) Create(ctx context.Context, requestData models.CreateUserRequest) (userId *models.UserId, err error) {
	userId = &models.UserId{}

	q := `
INSERT INTO users (username)
VALUES ($1)
RETURNING users.id;
		`
	err = ur.postgresConnection.QueryRow(ctx, q, requestData.Username).Scan(&(*userId).Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return nil, apperror.UsernameAlreadyExists
		}
		return nil, err
	}
	return userId, nil
}

func (ur *userRepository) DoesUserIdExist(ctx context.Context, userId int) (bool, error) {
	q := `
SELECT id
FROM users
WHERE id = $1;`
	rows, err := ur.postgresConnection.Query(ctx, q, userId)

	if err != nil {
		return false, err
	}
	if !rows.Next() {
		return false, nil
	}
	rows.Close()
	return true, nil
}
