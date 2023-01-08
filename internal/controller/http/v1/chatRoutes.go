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

type chatRoutes struct {
	chatUseCase *usecases.ChatUseCase
	logger      logger.Interface
}

func newChatRoutes(handler *gin.RouterGroup, chatUseCase *usecases.ChatUseCase, logger logger.Interface) {
	uC := &chatRoutes{
		chatUseCase: chatUseCase,
		logger:      logger,
	}

	chatHandlerGroup := handler.Group("/chats")
	{
		chatHandlerGroup.POST("/add", uC.createChat)
	}
}

type chatId struct {
	Id int `json:"chatId"`
}

// @Summary     Create new chat
// @Description Create a new user with name and users.
// @ID          createChat
// @Tags  	    chat
// @Accept      json
// @Produce     json
// @Success     200 {object} chatId
// @Failure	    400 {object} errorJSON
// @Failure	    500 {object} errorJSON
// @Router      /chats/add [post]
func (r *chatRoutes) createChat(c *gin.Context) {
	chatDTO := entities.ChatDTO{}

	err := c.BindJSON(&chatDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorJSON{err.Error()})
		return
	}

	newChatID, err := r.chatUseCase.CreateChat(chatDTO)
	if err != nil {
		if errors.Is(err, apperror.UserIDNotFound) {
			c.JSON(http.StatusBadRequest, errorJSON{err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, errorJSON{"Internal Server Error"})
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusCreated, chatId{newChatID})
}
