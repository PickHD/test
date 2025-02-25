package service

import (
	"context"
	"test/internal/config"
	"test/internal/model"
	"test/internal/repository"
)

type (
	AnimalService interface {
		Create(ctx context.Context, req *model.CreateAnimalRequest) (*model.Animal, error)
		GetAll(ctx context.Context) (*[]model.Animal, error)
	}

	AnimalServiceImpl struct {
		Context    context.Context
		Config     *config.Configuration
		AnimalRepo repository.AnimalRepository
	}
)

func NewAnimalService(ctx context.Context, config *config.Configuration, animalRepo repository.AnimalRepository) *AnimalServiceImpl {
	return &AnimalServiceImpl{
		Context:    ctx,
		Config:     config,
		AnimalRepo: animalRepo,
	}
}

func (as *AnimalServiceImpl) Create(ctx context.Context, req *model.CreateAnimalRequest) (*model.Animal, error) {
	return as.AnimalRepo.Create(ctx, req)
}

func (as *AnimalServiceImpl) GetAll(ctx context.Context) (*[]model.Animal, error) {
	return as.AnimalRepo.GetAll(ctx)
}
