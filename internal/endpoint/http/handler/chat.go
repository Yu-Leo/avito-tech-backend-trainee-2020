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
	chatService *service.ChatService
	logger      logger.Interface
}

func NewChatRoutes(handler *gin.RouterGroup, chatService *service.ChatService, logger logger.Interface) {
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

type chatId struct {
	Id int `json:"chatId"`
}

// CreateChat
// @Summary     Create new chat
// @Description Create a new chat with name and users.
// @ID          createChat
// @Tags  	    chat
// @Accept      json
// @Produce     json
// @Success     200 {object} chatId
// @Failure	    400 {object} errorJSON
// @Failure	    500 {object} errorJSON
// @Router      /chats/add [post]
func (r *chatRoutes) CreateChat(c *gin.Context) {
	chatDTO := models.CreateChatDTO{}

	err := c.BindJSON(&chatDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorJSON{err.Error()})
		return
	}

	newChatID, err := r.chatService.CreateChat(chatDTO)
	if err != nil {
		if errors.Is(err, apperror.IDNotFound) {
			c.JSON(http.StatusBadRequest, errorJSON{err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, errorJSON{"Internal Server Error"})
		r.logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusCreated, chatId{newChatID})
}

// GetUserChats
// @Summary     Get a list of user chats by ID
// @Description Get a list of user chats by ID.
// @ID          getUserChats
// @Tags  	    chat
// @Accept      json
// @Produce     json
// @Success     200 {list} models.GetUserChatsDTOAnswer
// @Failure	    400 {object} errorJSON
// @Failure	    500 {object} errorJSON
// @Router      /chats/get [post]
func (r *chatRoutes) GetUserChats(c *gin.Context) {
	userChatsDTORequest := models.GetUserChatsDTORequest{}

	err := c.BindJSON(&userChatsDTORequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorJSON{err.Error()})
		return
	}

	userChatsDTOAnswer, err := r.chatService.GetUserChats(userChatsDTORequest)
	if err != nil {
		if errors.Is(err, apperror.IDNotFound) {
			c.JSON(http.StatusBadRequest, errorJSON{err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, errorJSON{"Internal Server Error"})
		r.logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusCreated, userChatsDTOAnswer)
}
