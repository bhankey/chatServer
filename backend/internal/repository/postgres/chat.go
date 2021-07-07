package postgres

import (
	"chatServer/internal/models"
	"chatServer/pkg/db/sqlstore"
	"time"
)

type ChatRepo struct {
	db *sqlstore.Store
}

func NewChatRepo(db *sqlstore.Store) *ChatRepo {
	return &ChatRepo{
		db: db,
	}
}

func (r *ChatRepo) Create(name string, users []models.User) (id int, err error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var chatId int
	if err := tx.QueryRowx("INSERT INTO chat (name, created_at) VALUES ($1, $2) RETURNING id", name, time.Now()).Scan(&chatId); err != nil {
		return 0, err
	}
	for _, user := range users {
		if err := tx.QueryRowx("INSERT INTO chat_users (user_id, chat_id) VALUES ($1, $2)", user.Id, chatId).Err(); err != nil {
			return 0, err
		}
	}
	err = tx.Commit()
	return chatId, err
}

func (r *ChatRepo) GetById(chatId int) (chat *models.Chat, err error) {
	chat = &models.Chat{}
	row := r.db.QueryRowx("SELECT * FROM chat WHERE id = $1", chatId)
	if err := row.StructScan(&chat); err != nil {
		return nil, err
	}
	return chat, nil
}

func (r *ChatRepo) GetByUserId(userId int) (chats []models.Chat, err error) {
	rows, err := r.db.Queryx("SELECT Distinct chat_id FROM chat_users INNER JOIN users u on chat_users.user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// TODO in progress
	chats = make([]models.Chat, 0)
	var intrSlice []interface{}
	for rows.Next() {
		var chat models.Chat
		var tmp []interface{}
		if tmp, err = rows.SliceScan(); err != nil {
			return nil, err
		}
		intrSlice = append(intrSlice, tmp...)
		chats = append(chats, chat)
	}
	chatIds := make([]int, len(intrSlice))
	for i := range intrSlice {
		chatIds[i] = int(int(intrSlice[i].(int64)))
	}

	return chats, nil
}
