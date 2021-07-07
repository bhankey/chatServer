package models

import "time"

type Chat struct {
	Id        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UsersId   []int64   `db:"usersid" json:"users_id"`
}

type AddChat struct {
	Name   string  `json:"name"`
	UserId []int64 `json:"users"`
}

type ChatId struct {
	ChatId int `json:"chat_id"`
}

type ChatIdS struct {
	ChatId int `json:"chat"`
}
