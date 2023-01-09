package http

import (
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/endpoint/http/handler"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/service"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Yu-Leo/avito-tech-backend-trainee-2020/docs"
)

// @title           Avito-tech backend trainee task 2020
// @version         1.0

// @contact.name   Lev Yuvensky
// @contact.email  levayu22@gmail.com

// @host      127.0.0.1:9000
// @BasePath  /

func NewRouter(ginEngine *gin.Engine, logger logger.Interface,
	userUseCase *service.UserService, chatUseCase *service.ChatService) {

	// Routers
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router := ginEngine.Group("")
	{
		handler.NewUserRoutes(router, userUseCase, logger)
		handler.NewChatRoutes(router, chatUseCase, logger)
	}

}
