package repository

import (
	"context"
	"errors"
	"test/internal/config"
	"test/internal/model"

	"gorm.io/gorm"
)

type (
	AuthRepository interface {
		Create(ctx context.Context, req *model.RegisterRequest) (*model.User, error)
		VerifyUser(ctx context.Context, req *model.LoginRequest) (*model.User, error)
	}

	AuthRepostoryImpl struct {
		Context context.Context
		Config  *config.Configuration
		DB      *gorm.DB
	}
)

func NewAuthRepository(ctx context.Context, config *config.Configuration, db *gorm.DB) *AuthRepostoryImpl {
	return &AuthRepostoryImpl{
		Context: ctx,
		Config:  config,
		DB:      db,
	}
}

func (ar *AuthRepostoryImpl) Create(ctx context.Context, req *model.RegisterRequest) (*model.User, error) {
	user := model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := ar.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ar *AuthRepostoryImpl) VerifyUser(ctx context.Context, req *model.LoginRequest) (*model.User, error) {
	var user model.User

	err := ar.DB.Model(&model.User{}).Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
	}

	return &user, nil
}
