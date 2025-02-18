package service

import (
	"context"
	"test/internal/config"
	"test/internal/model"
	"test/internal/repository"
)

type (
	AuthService interface {
		Register(ctx context.Context, req *model.RegisterRequest) (*model.User, error)
		Login(ctx context.Context, req *model.LoginRequest) (string, error)
	}

	AuthServiceImpl struct {
		Context  context.Context
		Config   *config.Configuration
		AuthRepo repository.AuthRepository
	}
)

func NewAuthService(ctx context.Context, config *config.Configuration, authRepo repository.AuthRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		Context:  ctx,
		Config:   config,
		AuthRepo: authRepo,
	}
}

func (as *AuthServiceImpl) Register(ctx context.Context, req *model.RegisterRequest) (*model.User, error) {
	result, err := as.AuthRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (as *AuthServiceImpl) Login(ctx context.Context, req *model.LoginRequest) (string, error) {
	user, err := as.AuthRepo.VerifyUser(ctx, req)
	if err != nil {
		return "", err
	}

	//TODO : generate jwt here

	return "", nil
}
