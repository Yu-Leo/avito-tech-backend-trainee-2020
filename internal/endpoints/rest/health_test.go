package rest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/config"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/repositories/psql"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/services"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/httpserver"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/logger"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/postgresql"
)

func TestHealthEndpoint(t *testing.T) {
	cfg := &config.Config{
		Server: config.ServerConfig{
			Host: "0.0.0.0",
			Port: 9000,
		},
		Logger: config.LoggerConfig{
			Level: "debug",
		},
		Storage: config.StorageConfig{
			Host:     "localhost",
			Port:     5432,
			Database: "postgres",
			Username: "postgres",
			Password: "postgres",
		},
	}

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

	userService := services.NewUserService(userRepository)
	chatService := services.NewChatService(chatRepository)
	messageService := services.NewMessageService(messageRepository)

	ginEngine := gin.Default()
	NewRouter(ginEngine, l, userService, chatService, messageService)
	httpserver.New(ginEngine, httpserver.HostPort(cfg.Server.Host, cfg.Server.Port))

	req, _ := http.NewRequest("GET", "http://localhost/health", nil)
	w := httptest.NewRecorder()
	ginEngine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
