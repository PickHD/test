package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	Animal struct {
		ID        uint
		Name      string
		Type      string
		Color     string
		CreatedAt time.Time
	}

	CreateAnimalRequest struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Color string `json:"color"`
	}
)

func (v CreateAnimalRequest) Validate() error {
	if err := validation.ValidateStruct(&v,
		validation.Field(&v.Name, validation.Required),
		validation.Field(&v.Type, validation.Required),
		validation.Field(&v.Color, validation.Required)); err != nil {
		return err
	}

	return nil
}
