package app

import (
	"context"
	"fmt"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/config"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/endpoint/http"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repository/psql"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/service"
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
	defer func() {
		postgresConnection.Release()
		l.Info("Close Postgres connection")
	}()

	l.Info("Open Postgres connection")

	userRepository := psql.NewPostgresUserRepository(postgresConnection)
	chatRepository := psql.NewPostgresChatRepository(postgresConnection)
	messageRepository := psql.NewPostgresMessageRepository(postgresConnection)

	userService := service.NewUserService(userRepository)
	chatService := service.NewChatService(chatRepository)
	messageService := service.NewMessageService(messageRepository)

	ginEngine := gin.Default()
	http.NewRouter(ginEngine, l, userService, chatService, messageService)
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
