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
	AnimalController interface {
		Create(ctx *fiber.Ctx) error
		GetAll(ctx *fiber.Ctx) error
	}

	AnimalControllerImpl struct {
		Context   context.Context
		Config    *config.Configuration
		AnimalSvc service.AnimalService
	}
)

func NewAnimalController(ctx context.Context, config *config.Configuration, animalSvc service.AnimalService) *AnimalControllerImpl {
	return &AnimalControllerImpl{
		Context:   ctx,
		Config:    config,
		AnimalSvc: animalSvc,
	}
}

func (ac *AnimalControllerImpl) Create(ctx *fiber.Ctx) error {
	var req *model.CreateAnimalRequest

	if err := ctx.BodyParser(&req); err != nil {
		return helper.NewResponses[any](ctx, fiber.StatusBadRequest, "failed parse request", nil, err)
	}

	if err := req.Validate(); err != nil {
		return helper.NewResponses[any](ctx, fiber.StatusBadRequest, "failed validate request", nil, err)
	}

	result, err := ac.AnimalSvc.Create(ac.Context, req)
	if err != nil {
		return helper.NewResponses[any](ctx, fiber.StatusInternalServerError, "failed create animal", nil, err)
	}

	return helper.NewResponses[any](ctx, fiber.StatusCreated, "success create animal", result, nil)
}

func (ac *AnimalControllerImpl) GetAll(ctx *fiber.Ctx) error {
	result, err := ac.AnimalSvc.GetAll(ac.Context)
	if err != nil {
		return helper.NewResponses[any](ctx, fiber.StatusInternalServerError, "failed get all animal", nil, err)
	}

	return helper.NewResponses[any](ctx, fiber.StatusOK, "success get all animal", result, nil)
}
