package sqlstore

import (
	"github.com/jmoiron/sqlx"
)

type Store struct {
	*sqlx.DB
}

func (s *Store) Close() error {
	return s.DB.Close()
}
