package end_to_end_tests

import "time"

// Users

type CreateUserResponse struct {
	Id int `json:"userId"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
}

// Chats

type CreateChatRequest struct {
	Name  string `json:"name" binding:"required"`
	Users []int  `json:"users" binding:"required"`
}

type CreateChatResponse struct {
	Id int `json:"chatId"`
}

type GetUserChatsRequest struct {
	User int `json:"user" binding:"required"`
}

type GetUserChatsResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Users     []int     `json:"users"`
	CreatedAt time.Time `json:"createdAt"`
}

// Messages

type CreateMessageRequest struct {
	ChatId int    `json:"chat" binding:"required"`
	UserId int    `json:"author" binding:"required"`
	Text   string `json:"text" binding:"required"`
}

type CreateMessageResponse struct {
	Id int `json:"messageId"`
}

type GetChatMessagesRequest struct {
	ChatId int `json:"chat" binding:"required"`
}

type GetChatMessagesResponse struct {
	Id        int       `json:"id"`
	ChatId    int       `json:"chat"`
	UserId    int       `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}
