package service

import (
	"chatServer/internal/models"
	"chatServer/internal/repository"
)

type User interface {
	Create(userName string) (id int, err error)
	GetById(userId int) (user *models.User, err error)
}

type Chat interface {
	Create(name string, users []int) (id int, err error)
	GetById(chatId int) (chat *models.Chat, err error)
	GetByUserId(userId int) (chats []models.Chat, err error)
	//	AddUsers(chatId int, users []User) err error
}

type Message interface {
	Create(chatId int, userId int, text string) (id int, err error)
	GetByChatId(chatId int) (messages []models.Message, err error)
}

type Service struct {
	User    User
	Chat    Chat
	Message Message
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		User:    NewUserService(repository.User),
		Chat:    NewChatService(repository.Chat),
		Message: NewMessageService(repository.Message),
	}
}
