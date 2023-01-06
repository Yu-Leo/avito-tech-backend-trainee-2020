package v1

import (
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/usecases"
	"github.com/gin-gonic/gin"
)

func NewRouter(ginEngine *gin.Engine, userUseCase *usecases.UserUseCase) {
	// Routers
	router := ginEngine.Group("/v1")
	{
		newUserRoutes(router, userUseCase)
	}
}
