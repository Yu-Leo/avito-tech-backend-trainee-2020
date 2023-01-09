package models

import "time"

type Message struct {
	Id        int       `json:"id"`
	ChatId    int       `json:"chat"`
	UserId    int       `json:"author"`
	Text      int       `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}
