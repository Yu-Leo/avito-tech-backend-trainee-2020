package v1

import (
	"errors"
	"fmt"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/apperror"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/entities"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/usecases"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userRoutes struct {
	userUseCase *usecases.UserUseCase
	logger      logger.Interface
}

type errorJSON struct {
	Message string `json:"message"`
}

func newUserRoutes(handler *gin.RouterGroup, userUseCase *usecases.UserUseCase, logger logger.Interface) {
	uR := &userRoutes{
		userUseCase: userUseCase,
		logger:      logger,
	}

	userHandlerGroup := handler.Group("/users")
	{
		userHandlerGroup.POST("", uR.createUser)
	}
}

func (r *userRoutes) createUser(c *gin.Context) {
	userDTO := entities.UserDTO{}

	err := c.BindJSON(&userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorJSON{err.Error()})
		return
	}
	fmt.Println(userDTO)

	userid, err := r.userUseCase.CreateUser(userDTO)
	if err != nil {
		if errors.Is(err, apperror.UsernameAlreadyExists) {
			c.JSON(http.StatusBadRequest, errorJSON{err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, errorJSON{"Internal Server Error"})
		r.logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusCreated,
		struct {
			Id int `json:"userId"`
		}{userid})
}
