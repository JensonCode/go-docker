package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
