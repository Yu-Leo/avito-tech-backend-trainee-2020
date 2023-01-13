package models

import "time"

type Chat struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Users     []User    `json:"users"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateChatDTO struct {
	Name  string `json:"name" binding:"required"`
	Users []int  `json:"users" binding:"required"`
}

type GetUserChatsDTORequest struct {
	User int `json:"user" binding:"required"`
}

type GetUserChatsDTOAnswer struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Users     []int     `json:"users"`
	CreatedAt time.Time `json:"createdAt"`
}

type ChatId struct {
	Id int `json:"chatId"`
}
