package service

import (
	"vk-task/internal/models"
	"vk-task/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (models.User, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (string, error)
}

type Advertisement interface {
	Create(login string, input models.Advert) (models.Advert, error)
}

type Service struct {
	Authorization
	Advertisement
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		Advertisement: NewAdvertService(repo),
	}
}
