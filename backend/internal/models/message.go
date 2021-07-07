package models

import "time"

type Message struct {
	Id        int       `db:"id" json:"id"`
	Chat      int       `db:"chatid" json:"chat"`
	Author    int       `db:"author" json:"author"`
	Text      string    `db:"text" json:"text"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type AddMessage struct {
	ChatId int    `json:"chat"`
	UserId int    `json:"author"`
	Text   string `json:"text"`
}

type MessageId struct {
	MessageId int `json:"message_id"`
}
