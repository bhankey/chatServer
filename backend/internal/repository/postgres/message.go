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
	messagesRows, err := r.db.Queryx("SELECT message.id, message.chatid, message.userid as author, message.text, message.created_at FROM message WHERE message.chatid = $1 ORDER BY message.created_at DESC", chatId)
	if err != nil {
		return nil, err
	}
	defer messagesRows.Close()

	messages = make([]models.Message, 0)

	for messagesRows.Next() {
		var massage models.Message
		if err := messagesRows.StructScan(&massage); err != nil {
			return nil, err
		}
		messages = append(messages, massage)
	}
	return messages, nil
}
