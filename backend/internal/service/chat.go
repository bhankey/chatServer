package service

import (
	"chatServer/internal/models"
	"chatServer/internal/repository"
)

type ChatService struct {
	repo repository.Chat
}

func NewChatService(repo repository.Chat) *ChatService {
	return &ChatService{
		repo: repo,
	}
}

func (s *ChatService) Create(name string, usersId []int) (id int, err error) {
	users := make([]models.User, len(usersId))
	for i, id := range usersId {
		users[i] = models.User{
			Id: id,
		}
	}
	return s.repo.Create(name, users)
}

func (s *ChatService) GetById(chatId int) (chat *models.Chat, err error) {
	return s.repo.GetById(chatId)
}

func (s *ChatService) GetByUserId(userId int) (chats []models.Chat, err error) {
	return s.repo.GetByUserId(userId)
}

//type Chat interface {
//	Create(name string, users []models.User) (id int, err error)
//	GetById(chatId int) (chat *models.Chat, err error)
//	GetByUserId(userId int) (chats []models.Chat, err error)
//	//	AddUsers(chatId int, users []User) err error
//}
