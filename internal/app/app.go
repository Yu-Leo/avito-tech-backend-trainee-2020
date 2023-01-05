package app

import (
	"context"
	"fmt"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/config"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/postgresql"
	"github.com/jackc/pgx/v5/pgxpool"
)

func testCreateUser(pc *pgxpool.Pool) {
	q := `
		INSERT INTO users (username)
		VALUES ($1)
		RETURNING users.id;
			`
	var userID int
	ctx := context.Background()
	err := pc.QueryRow(ctx, q, "username 9").Scan(&userID)

	if err != nil {
		panic(err)
	}
	fmt.Println("New user id:", userID)

}

func Run(cfg *config.Config) {
	postgresClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)

	if err != nil {
		panic(err)
	}
	defer postgresClient.Close()

	testCreateUser(postgresClient)
}
