package repository

import (
	"github.com/jmoiron/sqlx"
	"vk-task/internal/models"
)

type Authorization interface {
	CreateUser(login models.User) (models.User, error)
	GetUser(login, password string) (models.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
