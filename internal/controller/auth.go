package controller

import (
	"context"
	"test/internal/config"
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

func NewAuthController(ctx context.Context, config *config.Configuration, authSvc service.AuthService)
