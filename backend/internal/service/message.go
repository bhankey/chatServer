package service

import (
	"chatServer/internal/models"
	"chatServer/internal/repository"
)

type MessageService struct {
	repo repository.Message
}

//type Message interface {
//	Create(chatId int, userId int, text string) (id int, err error)
//	GetByChatId(chatId int) (messages []models.Message, err error)
//}

func NewMessageService(repo repository.Message) *MessageService {
	return &MessageService{
		repo: repo,
	}
}

func (s *MessageService) Create(chatId int, userId int, text string) (id int, err error) {
	return s.repo.Create(chatId, userId, text)
}

func (s *MessageService) GetByChatId(chatId int) (messages []models.Message, err error) {
	return s.repo.GetByChatId(chatId)
}
