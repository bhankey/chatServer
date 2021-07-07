package postgres

import (
	"chatServer/internal/models"
	"chatServer/pkg/db/sqlstore"
	"time"
)

type MessageRepo struct {
	db *sqlstore.Store
}

func NewMessageRepo(db *sqlstore.Store) *MessageRepo {
	return &MessageRepo{
		db: db,
	}
}

func (r *MessageRepo) Create(chatId int, userId int, text string) (id int, err error) {
	var messageId int
	if err := r.db.QueryRowx(
		"INSERT INTO message (userid, chatid, text, created_at) VALUES ($1, $2, $3, $4) RETURNING id", userId, chatId, text, time.Now()).
		Scan(&messageId); err != nil {
		return 0, err
	}
	return messageId, nil
}

func (r *MessageRepo) GetByChatId(chatId int) (messages []models.Message, err error) {
	return nil, err
}
