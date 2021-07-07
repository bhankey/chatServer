package models

import "time"

type User struct {
	Id        int       `json:"id"`
	UserName  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}
