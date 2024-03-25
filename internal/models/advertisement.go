package models

import (
	"errors"
	"time"
	"vk-task/pkg/validation"
)

type Advert struct {
	Title       string    `json:"title" binding:"required"`
	Text        *string   `json:"text"`
	Image       *string   `json:"image"`
	Price       *float32  `json:"price"`
	PostingDate time.Time `json:"posting_date"`
	Owner       string    `json:"owner"`
}

func (advert Advert) Validate() error {
	err := ""
	if len(advert.Title) > 100 {
		err += "Длина текста не должна превышать 1000\n"
	}
	if advert.Text != nil && len(*advert.Text) > 1000 {
		err += "Длина текста не должна превышать 1000\n"
	}
	if advert.Text != nil && len(*advert.Text) > 1000 {
		err += "Длина текста не должна превышать 1000\n"
	}
	if advert.Image != nil {
		if errImage := validation.ValidImage(*advert.Image); errImage != nil {
			err += errImage.Error() + "\n"
		}
	}
	if advert.Price != nil && *advert.Price < 0 {
		err += "Цена должна быть положительной\n"
	}
	if err == "" {
		return nil
	}
	return errors.New(err)
}
