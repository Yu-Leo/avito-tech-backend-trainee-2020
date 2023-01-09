package models

import "time"

type ChatDTO struct {
	Name  string `json:"name"`
	Users []int  `json:"users"`
}

type Chat struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Users     []User    `json:"users"`
	CreatedAt time.Time `json:"createdAt"`
}
