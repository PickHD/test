package service

import (
	"context"
	"errors"
	"test/internal/config"
	"test/internal/helper"
	"test/internal/model"
	"test/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	hashedPass, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	req.Password = hashedPass

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

	if isValid := helper.VerifyPassword(req.Password, user.Password); !isValid {
		return "", errors.New("invalid credential")
	}

	token, err := as.generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (as *AuthServiceImpl) generateJWT(user *model.User) (string, error) {
	var JWTExpire = time.Duration(as.Config.Jwt.Expire) * time.Hour

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(JWTExpire).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signToken, err := token.SignedString([]byte(as.Config.Jwt.Secret))
	if err != nil {
		return "", err
	}

	return signToken, nil
}
