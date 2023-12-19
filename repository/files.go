package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type (
	FilesRepository interface {
		CreateFiles(ctx context.Context, files entities.Files) (entities.Files, error)
	}

	filesRepository struct {
		db *gorm.DB
	}
)

func NewFilesRepository(db *gorm.DB) FilesRepository {
	return &filesRepository{
		db: db,
	}
}

func (r *filesRepository) CreateFiles(ctx context.Context, files entities.Files) (entities.Files, error) {
	if err := r.db.Create(&files).Error; err != nil {
		return entities.Files{}, err
	}

	return files, nil
}
