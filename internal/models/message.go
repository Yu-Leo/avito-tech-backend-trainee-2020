package models

import "time"

type CreateMessageDTO struct {
	ChatId int    `json:"chat"`
	UserId int    `json:"author"`
	Text   string `json:"text"`
}

type Message struct {
	Id        int       `json:"id"`
	ChatId    int       `json:"chat"`
	UserId    int       `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}
