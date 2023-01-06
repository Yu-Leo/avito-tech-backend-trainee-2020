package app

import (
	"context"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/config"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/controller/http/v1"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repositories"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/usecases"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/httpserver"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/postgresql"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	postgresClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		panic(err)
	}
	defer postgresClient.Close()

	userRepository := repositories.NewPostgresUserRepository(postgresClient)

	userUserCase := usecases.NewUserUserCase(userRepository)

	ginEngine := gin.Default()

	v1.NewRouter(ginEngine, userUserCase)

	httpserver.New(ginEngine, httpserver.Port(cfg.Server.Port))
	ginEngine.Run()
}
