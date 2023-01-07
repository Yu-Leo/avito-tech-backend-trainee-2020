package app

import (
	"context"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/config"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/controller/http/v1"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repositories"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/usecases"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/httpserver"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/logger"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/postgresql"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func Run(cfg *config.Config) {
	l := logger.NewLogger(cfg.Logger.Level)

	l.Info("Run application")

	postgresClient, err := postgresql.NewClient(context.TODO(), 2, cfg.Storage)
	if err != nil {
		l.Fatal(err.Error())
	}
	l.Info("Open Postgres connection")

	defer func(postgresClient *pgx.Conn, ctx context.Context) {
		err := postgresClient.Close(ctx)
		if err == nil {
			l.Info("Close Postgres connection")
		} else {
			l.Error(err.Error())
		}
	}(postgresClient, context.Background())

	userRepository := repositories.NewPostgresUserRepository(postgresClient)
	chatRepository := repositories.NewPostgresChatRepository(postgresClient)

	userUserCase := usecases.NewUserUserCase(userRepository)
	chatUserCase := usecases.NewChatUserCase(chatRepository)

	ginEngine := gin.Default()

	v1.NewRouter(ginEngine, l, userUserCase, chatUserCase)

	httpserver.New(ginEngine, httpserver.Port(cfg.Server.Port))
	ginEngine.Run()
}
