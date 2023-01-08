package v1

import (
	"errors"
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
		userHandlerGroup.POST("/add", uR.createUser)
	}
}

type userId struct {
	Id int `json:"userId"`
}

// @Summary     Create new user
// @Description Create a new user with username.
// @ID          createUser
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Success     200 {object} userId
// @Failure	    400 {object} errorJSON
// @Failure	    500 {object} errorJSON
// @Router      /users/add [post]
func (r *userRoutes) createUser(c *gin.Context) {
	userDTO := entities.UserDTO{}

	err := c.BindJSON(&userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorJSON{err.Error()})
		return
	}

	newUserId, err := r.userUseCase.CreateUser(userDTO)
	if err != nil {
		if errors.Is(err, apperror.UsernameAlreadyExists) {
			c.JSON(http.StatusBadRequest, errorJSON{err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, errorJSON{"Internal Server Error"})
		r.logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusCreated, userId{newUserId})
}
