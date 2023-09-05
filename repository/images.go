package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/Caknoooo/golang-clean_template/entities"
)

type ImageRepository interface {
}

type imageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) *imageRepository {
	return &imageRepository{
		db: db,
	}
}

func (r *imageRepository) Create(ctx context.Context, image entities.Image) (entities.Image, error) {
	if err := r.db.WithContext(ctx).Create(&image).Error; err != nil {
		return entities.Image{}, err
	}

	return image, nil
}
