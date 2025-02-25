package controller

import (
	"context"
	"test/internal/config"
	"test/internal/helper"
	"test/internal/model"
	"test/internal/service"

	"github.com/gofiber/fiber/v2"
)

type (
	AuthController interface {
		Register(ctx *fiber.Ctx) error
		Login(ctx *fiber.Ctx) error
	}

	AuthControllerImpl struct {
		Context context.Context
		Config  *config.Configuration
		AuthSvc service.AuthService
	}
)

func NewAuthController(ctx context.Context, config *config.Configuration, authSvc service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		Context: ctx,
		Config:  config,
		AuthSvc: authSvc,
	}
}

func (ac *AuthControllerImpl) Register(ctx *fiber.Ctx) error {
	var req *model.RegisterRequest

	if err := ctx.BodyParser(&req); err != nil {
		return helper.NewResponses[any](ctx, fiber.StatusBadRequest, "failed parse request", nil, err)
	}

	if err := req.Validate(); err != nil {
		return helper.NewResponses[any](ctx, fiber.StatusBadRequest, "failed validate request", nil, err)
	}

	result, err := ac.AuthSvc.Register(ac.Context, req)
	if err != nil {
		return helper.NewResponses[any](ctx, fiber.StatusInternalServerError, "failed register user", nil, err)
	}

	return helper.NewResponses[any](ctx, fiber.StatusCreated, "success register user", result, nil)
}

func (ac *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	var req *model.LoginRequest

	if err := ctx.BodyParser(&req); err != nil {
		return helper.NewResponses[any](ctx, fiber.StatusBadRequest, "failed parse request", nil, err)
	}

	if err := req.Validate(); err != nil {
		return helper.NewResponses[any](ctx, fiber.StatusBadRequest, "failed validate request", nil, err)
	}

	result, err := ac.AuthSvc.Login(ac.Context, req)
	if err != nil {
		return helper.NewResponses[any](ctx, fiber.StatusInternalServerError, "failed login user", nil, err)
	}

	return helper.NewResponses[any](ctx, fiber.StatusOK, "success login user", result, nil)
}
