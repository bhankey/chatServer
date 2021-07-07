package postgres

import (
	"chatServer/internal/models"
	"chatServer/pkg/db/sqlstore"
	"github.com/lib/pq"
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
	if err := row.Scan(&chat.Id, &chat.Name, &chat.CreatedAt, pq.Array(&chat.UsersId)); err != nil {
		return nil, err
	}
	return chat, nil
}

func (r *ChatRepo) GetByUserId(userId int) (chats []models.Chat, err error) {
	chatIdRows, err := r.db.Queryx("WITH chats AS (SELECT DISTINCT chat.id, chat.name, chat.created_at FROM chat INNER JOIN chat_users INNER JOIN users u ON chat_users.user_id = $1 ON chat.id = chat_users.chat_id), message_time AS (SELECT chat.id as chat_id, max(message.created_at) as latest_message_time FROM chat LEFT JOIN message ON chat.id = message.chatid GROUP BY chat.id), users_agg AS (SELECT chat.id as chat_id, array_agg(c.id) as users FROM chat INNER JOIN chat_users cu on chat.id = cu.chat_id INNER JOIN users c on cu.user_id = c.id GROUP BY chat.id) SELECT chats.id, chats.name, chats.created_at, users_agg.users as usersid FROM chats LEFT JOIN message_time ON chats.id = message_time.chat_id LEFT JOIN users_agg ON users_agg.chat_id = chats.id ORDER BY message_time.latest_message_time DESC NULLS LAST", userId)
	if err != nil {
		return nil, err
	}
	defer chatIdRows.Close()

	chats = make([]models.Chat, 0)

	for chatIdRows.Next() {
		var chat models.Chat
		if err := chatIdRows.Scan(&chat.Id, &chat.Name, &chat.CreatedAt, pq.Array(&chat.UsersId)); err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}

	return chats, nil
}
