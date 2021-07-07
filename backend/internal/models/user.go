package models

import "time"

type User struct {
	Id        int       `db:"id"`
	UserName  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
}

type UserId struct {
	UserId int `json:"user"`
}

type AddUser struct {
	UserName string `json:"username"`
}
