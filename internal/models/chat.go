package models

import "time"

type Chat struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Users     []User    `json:"users"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateChatRequest struct {
	Name  string `json:"name" binding:"required"`
	Users []int  `json:"users" binding:"required"`
}

type ChatId struct {
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


