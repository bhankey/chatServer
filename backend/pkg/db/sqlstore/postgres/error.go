package postgres

import (
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
)

func PgError(err error) error {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		err = fmt.Errorf(pgErr.Message)
		return err
	}
	return nil
}
