package repositories

import (
	"context"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/entities"
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
		return 0, err
	}
	return userID, err
}
