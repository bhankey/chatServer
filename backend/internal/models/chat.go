package models

import "time"

type Chat struct {
	Id        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	UserId    int       `db:"userid" json:"user_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
