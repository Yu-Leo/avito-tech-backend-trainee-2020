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

func (ur *userRepository) Create(ctx context.Context, user models.CreateUserRequest) (userId *models.UserId, err error) {
	userId = &models.UserId{}

	q := `
INSERT INTO users (username)
VALUES ($1)
RETURNING users.id;
		`
	err = ur.postgresConnection.QueryRow(ctx, q, user.Username).Scan(&(*userId).Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return nil, apperror.UsernameAlreadyExists
		}
		return nil, err
	}
	return userId, nil
}
