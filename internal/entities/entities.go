package entities

import "time"

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserDTO struct {
	Username string `json:"username"`
}

type Chat struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Users     []User    `json:"users"`
	CreatedAt time.Time `json:"createdAt"`
}

type Message struct {
	Id        int       `json:"id"`
	ChatId    int       `json:"chat"`
	UserId    int       `json:"author"`
	Text      int       `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}
