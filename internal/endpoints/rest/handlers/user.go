package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/apperror"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/services"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/logger"
)

type userRoutes struct {
	userService *services.UserService
	logger      logger.Interface
}

func NewUserRoutes(handler *gin.RouterGroup, userService *services.UserService, logger logger.Interface) {
	uR := &userRoutes{
		userService: userService,
		logger:      logger,
	}

	userHandlerGroup := handler.Group("/users")
	{
		userHandlerGroup.POST("/add", uR.CreateUser)
	}
}

// CreateUser
// @Summary     Create new user
// @Description Create a new user with username.
// @ID          createUser
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param createUserObject body models.CreateUserRequest true "Parameters for creating a user."
// @Success     201 {object} models.UserId
// @Failure	    400 {object} apperror.ErrorJSON
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /users/add [post]
func (r *userRoutes) CreateUser(c *gin.Context) {
	requestData := models.CreateUserRequest{}

	err := c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, apperror.ErrorJSON{
			Message:          apperror.ValidationErrorMsg,
			DeveloperMessage: err.Error()})
		return
	}

	newUserId, err := r.userService.CreateUser(requestData)
	if err != nil {
		if errors.Is(err, apperror.UsernameAlreadyExists) || errors.Is(err, apperror.TooLongName) {
			c.JSON(http.StatusBadRequest, apperror.ErrorJSON{
				Message:          apperror.ValidationErrorMsg,
				DeveloperMessage: err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, apperror.ErrorJSON{Message: apperror.InternalServerErrorMsg})
		r.logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusCreated, *newUserId)
}
