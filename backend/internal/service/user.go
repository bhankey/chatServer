package service

import (
	"chatServer/internal/models"
	"chatServer/internal/repository"
)

type UserService struct {
	repo repository.User
}

//type User interface {
//	Create(userName string) (id int, err error)
//	GetById(userId int) (user models.User, err error)
//}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(userName string) (id int, err error) {
	return s.repo.Create(userName)
}

func (s *UserService) GetById(userId int) (user *models.User, err error) {
	return s.repo.GetById(userId)
}
