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
)

func Run(cfg *config.Config) {
	l := logger.NewLogger(cfg.Logger.Level)

	l.Info("Run application")

	postgresClient, err := postgresql.NewClient(context.TODO(), 1, cfg.Storage)
	if err != nil {
		l.Fatal(err.Error())
	}

	l.Debug("Connected to postgres")

	defer postgresClient.Close()

	userRepository := repositories.NewPostgresUserRepository(postgresClient)

	userUserCase := usecases.NewUserUserCase(userRepository)

	ginEngine := gin.Default()

	v1.NewRouter(ginEngine, userUserCase, l)

	httpserver.New(ginEngine, httpserver.Port(cfg.Server.Port))
	ginEngine.Run()
}
