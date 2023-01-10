package models

import "time"

type CreateChatDTO struct {
	Name  string `json:"name"`
	Users []int  `json:"users"`
}

type GetUserChatsDTORequest struct {
	User int `json:"user"`
}

type GetUserChatsDTOAnswer struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Users     []int     `json:"users"`
	CreatedAt time.Time `json:"createdAt"`
}

type Chat struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Users     []User    `json:"users"`
	CreatedAt time.Time `json:"createdAt"`
}
