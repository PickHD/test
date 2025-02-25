package repository

import (
	"context"
	"test/internal/config"
	"test/internal/model"

	"gorm.io/gorm"
)

type (
	AnimalRepository interface {
		Create(ctx context.Context, req *model.CreateAnimalRequest) (*model.Animal, error)
		GetAll(ctx context.Context) (*[]model.Animal, error)
	}

	AnimalRepostoryImpl struct {
		Context context.Context
		Config  *config.Configuration
		DB      *gorm.DB
	}
)

func NewAnimalRepository(ctx context.Context, config *config.Configuration, db *gorm.DB) *AnimalRepostoryImpl {
	return &AnimalRepostoryImpl{
		Context: ctx,
		Config:  config,
		DB:      db,
	}
}

func (ar *AnimalRepostoryImpl) Create(ctx context.Context, req *model.CreateAnimalRequest) (*model.Animal, error) {
	animal := model.Animal{
		Name:  req.Name,
		Type:  req.Type,
		Color: req.Color,
	}

	err := ar.DB.Create(&animal).Error
	if err != nil {
		return nil, err
	}

	return &animal, nil
}

func (ar *AnimalRepostoryImpl) GetAll(ctx context.Context) (*[]model.Animal, error) {
	var animals []model.Animal

	if err := ar.DB.Model(&model.Animal{}).Find(&animals).Error; err != nil {
		return nil, err
	}

	return &animals, nil
}
