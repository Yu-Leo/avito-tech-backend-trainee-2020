package models

import "time"

type Message struct {
	Id        int       `json:"id"`
	ChatId    int       `json:"chat"`
	UserId    int       `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateMessageRequest struct {
	ChatId int    `json:"chat" binding:"required"`
	UserId int    `json:"author" binding:"required"`
	Text   string `json:"text" binding:"required"`
}

type GetChatMessagesRequest struct {
	ChatId int `json:"chat" binding:"required"`
}

type MessageId struct {
	Id int `json:"messageId"`
}
