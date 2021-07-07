package postgres

import (
	"chatServer/internal/models"
	"chatServer/pkg/db/sqlstore"
	"time"
)

type UserRepo struct {
	db *sqlstore.Store
}

func NewUserRepo(db *sqlstore.Store) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(userName string) (id int, err error) {
	if err := r.db.QueryRowx("INSERT INTO users (username, created_at) VALUES ($1, $2) RETURNING id", userName, time.Now()).Scan(&id); err != nil {
		return 0, err
	}
	return
}

func (r *UserRepo) GetById(userId int) (user *models.User, err error) {
	user = &models.User{}
	row := r.db.QueryRowx("SELECT * FROM users WHERE id = $1", userId)
	if err := row.StructScan(&user); err != nil {
		return nil, err
	}
	return user, nil
}
