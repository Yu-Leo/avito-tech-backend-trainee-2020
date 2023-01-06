package v1

import (
	"fmt"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/entities"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userRoutes struct {
	userUseCase *usecases.UserUseCase
}

func newUserRoutes(handler *gin.RouterGroup, userUseCase *usecases.UserUseCase) {
	uR := &userRoutes{userUseCase}

	userHandlerGroup := handler.Group("/users")
	{
		userHandlerGroup.POST("", uR.createUser)
	}
}

func (r *userRoutes) createUser(c *gin.Context) {
	userDTO := entities.UserDTO{}

	err := c.BindJSON(&userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			struct {
				Message string `json:"message"`
			}{err.Error()})
		return
	}
	fmt.Println(userDTO)

	userid, err := r.userUseCase.CreateUser(userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			struct {
				Message string `json:"message"`
			}{err.Error()})
		return
	}

	c.JSON(http.StatusOK,
		struct {
			Id int `json:"userId"`
		}{userid})
}
