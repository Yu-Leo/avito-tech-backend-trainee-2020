package handler

import (
	"errors"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/apperror"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/service"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type chatRoutes struct {
	chatUseCase *service.ChatService
	logger      logger.Interface
}

func NewChatRoutes(handler *gin.RouterGroup, chatUseCase *service.ChatService, logger logger.Interface) {
	uC := &chatRoutes{
		chatUseCase: chatUseCase,
		logger:      logger,
	}

	chatHandlerGroup := handler.Group("/chats")
	{
		chatHandlerGroup.POST("/add", uC.CreateChat)
	}
}

type chatId struct {
	Id int `json:"chatId"`
}

// CreateChat
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
func (r *chatRoutes) CreateChat(c *gin.Context) {
	chatDTO := models.ChatDTO{}

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
