package models

import "time"

type User struct {
	Id        int       `db:"id"`
	UserName  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
}
