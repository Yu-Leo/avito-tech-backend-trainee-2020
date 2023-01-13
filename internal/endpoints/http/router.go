package http

import (
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/endpoints/http/handlers"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/services"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Yu-Leo/avito-tech-backend-trainee-2020/docs"
)

// @title           Avito.tech's task for backend trainee (2020 year)
// @version         1.0

// @contact.name   Lev Yuvensky
// @contact.email  levayu22@gmail.com

// @host      127.0.0.1:9000
// @BasePath  /

func NewRouter(ginEngine *gin.Engine, logger logger.Interface,
	userService *services.UserService, chatService *services.ChatService, messageService *services.MessageService) {

	// Routers
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router := ginEngine.Group("")
	{
		handlers.NewUserRoutes(router, userService, logger)
		handlers.NewChatRoutes(router, chatService, logger)
		handlers.NewMessageRoutes(router, messageService, logger)
	}

}
