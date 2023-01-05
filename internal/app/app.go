package app

import (
	"context"
	"fmt"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/config"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/entities"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repositories"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/usecases"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/postgresql"
	"github.com/jackc/pgx/v5/pgxpool"
)

func testCreateUser(pc *pgxpool.Pool) {
	userRepository := repositories.NewPostgresUserRepository(pc)
	userUserCase := usecases.NewUserUserCase(userRepository)

	userDTO := entities.UserDTO{
		Username: "user 2",
	}

	userId, err := userUserCase.CreateUser(userDTO)
	if err != nil {
		panic(err)
	}
	fmt.Println("New user id", userId)
}

func Run(cfg *config.Config) {
	postgresClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)

	if err != nil {
		panic(err)
	}
	defer postgresClient.Close()

	testCreateUser(postgresClient)
}
