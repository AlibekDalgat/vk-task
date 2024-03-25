package service

import (
	"time"
	"vk-task/internal/models"
	"vk-task/internal/repository"
)

type AdvertService struct {
	repo repository.Advertisement
}

func NewAdvertService(repo repository.Advertisement) *AdvertService {

	return &AdvertService{repo}
}

func (adv *AdvertService) Create(login string, input models.Advert) (models.Advert, error) {
	if err := input.Validate(); err != nil {
		return models.Advert{}, err
	}
	input.Owner = login
	input.PostingDate = time.Now()
	return adv.repo.Create(input)
}
