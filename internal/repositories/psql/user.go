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

type userRepository struct {
	postgresConnection postgresql.Connection
}

func NewPostgresUserRepository(pc postgresql.Connection) repositories.UserRepository {
	return &userRepository{
		postgresConnection: pc,
	}
}

func (ur *userRepository) Create(ctx context.Context, user models.CreateUserDTO) (userId models.UserId, err error) {
	q := `
INSERT INTO users (username)
VALUES ($1)
RETURNING users.id;
		`
	err = ur.postgresConnection.QueryRow(ctx, q, user.Username).Scan(&userId.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return models.UserId{}, apperror.UsernameAlreadyExists
		}
		return models.UserId{}, err
	}
	return userId, nil
}
