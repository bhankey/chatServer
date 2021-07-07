package repository

import (
	"chatServer/internal/models"
	"chatServer/internal/repository/postgres"
	"chatServer/pkg/db/sqlstore"
)

type User interface {
	Create(userName string) (id int, err error)
	GetById(userId int) (user *models.User, err error)
}

type Chat interface {
	Create(name string, users []models.User) (id int, err error)
	GetById(chatId int) (chat *models.Chat, err error)
	GetByUserId(userId int) (chats []models.Chat, err error)
	//	AddUsers(chatId int, users []User) err error
}

type Message interface {
	Create(chatId int, userId int, text string) (id int, err error)
	GetByChatId(chatId int) (messages []models.Message, err error)
}

type Repository struct {
	User    User
	Chat    Chat
	Message Message
}

func NewRepository(db *sqlstore.Store) *Repository {
	return &Repository{
		User:    postgres.NewUserRepo(db),
		Chat:    postgres.NewChatRepo(db),
		Message: postgres.NewMessageRepo(db),
	}
}
