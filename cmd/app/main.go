package main

import (
	"context"
	"fmt"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/config"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/postgresql"
)

func main() {
	storageConfig := config.StorageConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "postgres",
		Username: "postgres",
		Password: "postgres",
	}
	postgreSQLClient, err := postgresql.NewClient(context.TODO(), 3, storageConfig)

	if err != nil {
		panic(err)
	}

	q := `
INSERT INTO users (username)
VALUES ($1)
RETURNING users.id;
	`
	var userID int
	ctx := context.Background()
	err = postgreSQLClient.QueryRow(ctx, q, "username 6").Scan(&userID)
	fmt.Println("New user id:", userID)

	if err != nil {
		panic(err)
	}

}
