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

type messageId struct {
	Id int `json:"messageId"`
}

// CreateMessage
// @Summary     Create new message
// @Description Create a new message with chat's name, author and text.
// @ID          createMessage
// @Tags  	    chat
// @Accept      json
// @Produce     json
// @Success     200 {object} messageId
// @Failure	    400 {object} errorJSON
// @Failure	    500 {object} errorJSON
// @Router      /messages/add [post]
func (r *messageRoutes) CreateMessage(c *gin.Context) {
	messageDTO := models.CreateMessageDTO{}

	err := c.BindJSON(&messageDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorJSON{err.Error()})
		return
	}
	newMessageID, err := r.messageService.CreateMessage(messageDTO)

	if err != nil {
		if errors.Is(err, apperror.IDNotFound) {
			c.JSON(http.StatusBadRequest, errorJSON{err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, errorJSON{apperror.InternalServerError.Error()})
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusCreated, messageId{newMessageID})
}

// GetChatMessages
// @Summary     Get chat messages
// @Description Get messages from chat.
// @ID          GetChatMessages
// @Tags  	    chat
// @Accept      json
// @Produce     json
// @Success     200 {list} models.Message
// @Failure	    400 {object} errorJSON
// @Failure	    500 {object} errorJSON
// @Router      /messages/add [post]
func (r *messageRoutes) GetChatMessages(c *gin.Context) {
	chatMessagesDTORequest := models.GetChatMessagesDRORequest{}

	err := c.BindJSON(&chatMessagesDTORequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorJSON{err.Error()})
		return
	}

	chatMessages, err := r.messageService.GetChatMessages(chatMessagesDTORequest)
	if err != nil {
		if errors.Is(err, apperror.IDNotFound) {
			c.JSON(http.StatusBadRequest, errorJSON{err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, errorJSON{apperror.InternalServerError.Error()})
		r.logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, *chatMessages)
}
