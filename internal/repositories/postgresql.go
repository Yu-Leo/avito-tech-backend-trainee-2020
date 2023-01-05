package repositories

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	postgresClient *pgxpool.Pool
}

func NewUserRepository(pc *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		postgresClient: pc,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, username string) (userID int, err error) {
	q := `
		INSERT INTO users (username)
		VALUES ($1)
		RETURNING users.id;
		`
	err = ur.postgresClient.QueryRow(ctx, q, username).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, err
}
