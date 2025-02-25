package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	User struct {
		ID        uint      `json:"id"`
		Email     string    `json:"email"`
		Password  string    `json:"-"`
		CreatedAt time.Time `json:"created_at"`
	}

	RegisterRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func (v RegisterRequest) Validate() error {
	if err := validation.ValidateStruct(&v,
		validation.Field(&v.Email, validation.Required),
		validation.Field(&v.Password, validation.Required)); err != nil {
		return err
	}

	return nil
}

func (v LoginRequest) Validate() error {
	if err := validation.ValidateStruct(&v,
		validation.Field(&v.Email, validation.Required),
		validation.Field(&v.Password, validation.Required)); err != nil {
		return err
	}

	return nil
}
