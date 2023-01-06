package v1

import (
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/usecases"
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

// @host      localhost:8080
// @BasePath  /api/v1

func NewRouter(ginEngine *gin.Engine, userUseCase *usecases.UserUseCase, logger logger.Interface) {
	// Routers
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router := ginEngine.Group("/v1")
	{
		newUserRoutes(router, userUseCase, logger)
	}

}
