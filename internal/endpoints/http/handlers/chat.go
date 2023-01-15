package handlers

import (
	"errors"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/apperror"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/models"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/internal/services"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type chatRoutes struct {
	chatService *services.ChatService
	logger      logger.Interface
}

func NewChatRoutes(handler *gin.RouterGroup, chatService *services.ChatService, logger logger.Interface) {
	uC := &chatRoutes{
		chatService: chatService,
		logger:      logger,
	}

	chatHandlerGroup := handler.Group("/chats")
	{
		chatHandlerGroup.POST("/add", uC.CreateChat)
		chatHandlerGroup.POST("/get", uC.GetUserChats)
	}
}

// CreateChat
// @Summary     Create new chat
// @Description Create a new chat with name and list of users.
// @ID          createChat
// @Tags  	    chat
// @Accept      json
// @Produce     json
// @Param createChatObject body models.CreateChatDTO true "Parameters for creating a chat."
// @Success     201 {object} models.ChatId
// @Failure	    400 {object} apperror.ErrorJSON
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /chats/add [post]
func (r *chatRoutes) CreateChat(c *gin.Context) {
	chatDTO := models.CreateChatDTO{}

	err := c.BindJSON(&chatDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, apperror.ErrorJSON{Message: err.Error()})
		return
	}

	newChatID, err := r.chatService.CreateChat(chatDTO)
	if err != nil {
		if errors.Is(err, apperror.IDNotFound) || errors.Is(err, apperror.ChatNameAlreadyExists) {
			c.JSON(http.StatusBadRequest, apperror.ErrorJSON{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, apperror.ErrorJSON{Message: apperror.InternalServerError.Error()})
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusCreated, *newChatID)
}

// GetUserChats
// @Summary     Get a list of user chats
// @Description Get a list of user chats by user ID.
// @ID          getUserChats
// @Tags  	    chat
// @Accept      json
// @Produce     json
// @Param getUserChatsObject body models.GetUserChatsDTORequest true "Parameters for getting user chats."
// @Success     200 {array} models.GetUserChatsDTOAnswer
// @Failure	    400 {object} apperror.ErrorJSON
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /chats/get [post]
func (r *chatRoutes) GetUserChats(c *gin.Context) {
	userChatsDTORequest := models.GetUserChatsDTORequest{}

	err := c.BindJSON(&userChatsDTORequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, apperror.ErrorJSON{Message: err.Error()})
		return
	}

	userChatsDTOAnswer, err := r.chatService.GetUserChats(userChatsDTORequest)
	if err != nil {
		if errors.Is(err, apperror.IDNotFound) {
			c.JSON(http.StatusBadRequest, apperror.ErrorJSON{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, apperror.ErrorJSON{Message: apperror.InternalServerError.Error()})
		r.logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, *userChatsDTOAnswer)
}
