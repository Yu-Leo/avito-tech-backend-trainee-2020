package repositories

import (
	"context"
	"errors"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/apperror"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/entities"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	postgresClient *pgxpool.Pool
}

func NewPostgresUserRepository(pc *pgxpool.Pool) UserRepository {
	return &userRepository{
		postgresClient: pc,
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, user entities.UserDTO) (userID int, err error) {
	q := `
		INSERT INTO users (username)
		VALUES ($1)
		RETURNING users.id;
		`
	err = ur.postgresClient.QueryRow(ctx, q, user.Username).Scan(&userID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return 0, apperror.UsernameAlreadyExists
		}
		return 0, err
	}
	return userID, nil
}
