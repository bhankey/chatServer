package service

import (
	"chatServer/internal/models"
	"chatServer/internal/repository"
	"fmt"
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
	if len(usersId) < 1 {
		return 0, fmt.Errorf("no users id was recived")
	}
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
