package app

import (
	"context"
	"fmt"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/config"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/controller/http/v1"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repositories"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/usecases"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/httpserver"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/logger"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/postgresql"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	l := logger.NewLogger(cfg.Logger.Level)

	l.Info("Run application")

	postgresConnection, err := postgresql.NewConnection(context.TODO(), 2, cfg.Storage)
	if err != nil {
		l.Fatal(err.Error())
	}
	l.Info("Open Postgres connection")

	userRepository := repositories.NewPostgresUserRepository(postgresConnection)
	chatRepository := repositories.NewPostgresChatRepository(postgresConnection)

	userUserCase := usecases.NewUserUserCase(userRepository)
	chatUserCase := usecases.NewChatUserCase(chatRepository)

	ginEngine := gin.Default()
	v1.NewRouter(ginEngine, l, userUserCase, chatUserCase)
	httpServer := httpserver.New(ginEngine, httpserver.HostPort(cfg.Server.Host, cfg.Server.Port))
	l.Info(fmt.Sprintf("Run server on %s:%d", cfg.Server.Host, cfg.Server.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info(fmt.Sprintf("Catch the %s signal", s.String()))
	case err = <-httpServer.Notify():
		l.Error(fmt.Sprintf("HTTPServer notify error: %e", err))
	}

	err = httpServer.Shutdown()
	if err == nil {
		l.Info("Shutdown HTTPServer")
	} else {
		l.Error(fmt.Sprintf("HTTPServer shutdown error: %e", err))
	}
}
