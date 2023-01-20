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

type messageRoutes struct {
	messageService *services.MessageService
	logger         logger.Interface
}

func NewMessageRoutes(handler *gin.RouterGroup, messageService *services.MessageService, logger logger.Interface) {
	uC := &messageRoutes{
		messageService: messageService,
		logger:         logger,
	}

	chatHandlerGroup := handler.Group("/messages")
	{
		chatHandlerGroup.POST("/add", uC.CreateMessage)
		chatHandlerGroup.POST("/get", uC.GetChatMessages)
	}
}

// CreateMessage
// @Summary     Create new message
// @Description Create a new message with chat's id, author's id and text.
// @ID          createMessage
// @Tags  	    message
// @Accept      json
// @Produce     json
// @Param createMessageObject body models.CreateMessageRequest true "Parameters for creating a message."
// @Success     201 {object} models.MessageId
// @Failure	    400 {object} apperror.ErrorJSON
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /messages/add [post]
func (r *messageRoutes) CreateMessage(c *gin.Context) {
	messageDTO := models.CreateMessageRequest{}

	err := c.BindJSON(&messageDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, apperror.ErrorJSON{Message: err.Error()})
		return
	}
	newMessageID, err := r.messageService.CreateMessage(messageDTO)

	if err != nil {
		if errors.Is(err, apperror.IDNotFound) || errors.Is(err, apperror.UserIsNotInChat) {
			c.JSON(http.StatusBadRequest, apperror.ErrorJSON{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, apperror.ErrorJSON{Message: apperror.InternalServerError.Error()})
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusCreated, *newMessageID)
}

// GetChatMessages
// @Summary     Get a list of chat messages
// @Description Get a list of chat messages by chat ID.
// @ID          GetChatMessages
// @Tags  	    message
// @Accept      json
// @Produce     json
// @Param getChatsMessagesObject body models.GetChatMessagesRequest true "Parameters for getting chat messages."
// @Success     200 {array} models.Message
// @Failure	    400 {object} apperror.ErrorJSON
// @Failure	    500 {object} apperror.ErrorJSON
// @Router      /messages/get [post]
func (r *messageRoutes) GetChatMessages(c *gin.Context) {
	chatMessagesDTORequest := models.GetChatMessagesRequest{}

	err := c.BindJSON(&chatMessagesDTORequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, apperror.ErrorJSON{Message: err.Error()})
		return
	}

	chatMessages, err := r.messageService.GetChatMessages(chatMessagesDTORequest)
	if err != nil {
		if errors.Is(err, apperror.IDNotFound) {
			c.JSON(http.StatusBadRequest, apperror.ErrorJSON{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, apperror.ErrorJSON{Message: apperror.InternalServerError.Error()})
		r.logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, *chatMessages)
}
