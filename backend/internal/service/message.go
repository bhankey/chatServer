package service

import (
	"chatServer/internal/models"
	"chatServer/internal/repository"
	"fmt"
)

type MessageService struct {
	messageRepo repository.Message
	chatRepo    repository.Chat
}

//type Message interface {
//	Create(chatId int, userId int, text string) (id int, err error)
//	GetByChatId(chatId int) (messages []models.Message, err error)
//}

func NewMessageService(message repository.Message, chat repository.Chat) *MessageService {
	return &MessageService{
		messageRepo: message,
		chatRepo:    chat,
	}
}

func (s *MessageService) Create(chatId int, userId int, text string) (int, error) {
	c, err := s.chatRepo.GetById(chatId)
	if err != nil {
		return 0, err
	}
	isAllowed := false
	for uId := range c.UsersId {
		if uId == userId {
			isAllowed = true
			break
		}
	}
	if !isAllowed {
		return 0, fmt.Errorf("user not allowed to write in this chat")
	}
	return s.messageRepo.Create(chatId, userId, text)
}

func (s *MessageService) GetByChatId(chatId int) (messages []models.Message, err error) {
	return s.messageRepo.GetByChatId(chatId)
}
