package postgres

import (
	"chatServer/pkg/db/sqlstore"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewStore(c *sqlstore.Config) (*sqlstore.Store, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)
	db, err := sqlx.Connect("pgx", psqlInfo)
	if err != nil {
		return nil, err
	}
	store := &sqlstore.Store{
		DB: db,
	}
	return store, nil
}
