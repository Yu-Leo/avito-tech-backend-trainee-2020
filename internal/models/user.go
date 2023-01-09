package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserDTO struct {
	Username string `json:"username" binding:"required"`
}
