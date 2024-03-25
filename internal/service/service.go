package service

import (
	"vk-task/internal/models"
	"vk-task/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (models.User, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (models.User, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
	}
}
