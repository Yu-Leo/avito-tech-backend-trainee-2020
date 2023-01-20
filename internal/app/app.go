package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/config"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/endpoints/rest"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repositories/psql"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/services"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/httpserver"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/logger"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/postgresql"
)

func Run(cfg *config.Config) {
	l := logger.NewLogger(cfg.Logger.Level)

	l.Info("Run application")

	postgresConnection, err := postgresql.NewConnection(context.Background(), 2, cfg.Storage)
	if err != nil {
		l.Fatal(err.Error())
	}
	defer func() {
		postgresConnection.Release()
		l.Info("Close Postgres connection")
	}()

	l.Info("Open Postgres connection")

	userRepository := psql.NewPostgresUserRepository(postgresConnection)
	chatRepository := psql.NewPostgresChatRepository(postgresConnection)
	messageRepository := psql.NewPostgresMessageRepository(postgresConnection)

	userService := services.NewUserService(userRepository)
	chatService := services.NewChatService(chatRepository)
	messageService := services.NewMessageService(messageRepository)

	ginEngine := gin.Default()
	rest.NewRouter(ginEngine, l, userService, chatService, messageService)
	httpServer := httpserver.New(ginEngine, cfg.Server.Host, cfg.Server.Port)
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
